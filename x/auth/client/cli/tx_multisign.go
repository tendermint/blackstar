package cli

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/version"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// GetSignCommand returns the sign command
func GetMultiSignCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "multisign [file] [name] [[signature]...]",
		Short: "Generate multisig signatures for transactions generated offline",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Sign transactions created with the --generate-only flag that require multisig signatures.

Read signature(s) from [signature] file(s), generate a multisig signature compliant to the
multisig key [name], and attach it to the transaction read from [file].

Example:
$ %s multisign transaction.json k1k2k3 k1sig.json k2sig.json k3sig.json

If the flag --signature-only flag is on, it outputs a JSON representation
of the generated signature only.

The --offline flag makes sure that the client will not reach out to an external node.
Thus account number or sequence number lookups will not be performed and it is
recommended to set such parameters manually.
`,
				version.AppName,
			),
		),
		RunE: makeMultiSignCmd(),
		Args: cobra.MinimumNArgs(3),
	}

	cmd.Flags().Bool(flagSigOnly, false, "Print only the generated signature, then exit")
	cmd.Flags().String(flags.FlagOutputDocument, "", "The document will be written to the given file instead of STDOUT")
	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flags.FlagChainID, "", "network chain ID")

	return cmd
}

func makeMultiSignCmd() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) (err error) {
		clientCtx := client.GetClientContextFromCmd(cmd)
		clientCtx, err = client.ReadTxCommandFlags(clientCtx, cmd.Flags())
		if err != nil {
			return err
		}
		cdc := clientCtx.Codec
		stdTx, err := authclient.ReadTxFromFile(clientCtx, args[0])
		if err != nil {
			return
		}

		backend, _ := cmd.Flags().GetString(flags.FlagKeyringBackend)

		inBuf := bufio.NewReader(cmd.InOrStdin())
		kb, err := keyring.New(sdk.KeyringServiceName(), backend, clientCtx.HomeDir, inBuf)
		if err != nil {
			return
		}

		multisigInfo, err := kb.Key(args[1])
		if err != nil {
			return
		}
		if multisigInfo.GetType() != keyring.TypeMulti {
			return fmt.Errorf("%q must be of type %s: %s", args[1], keyring.TypeMulti, multisigInfo.GetType())
		}

		multisigPub := multisigInfo.GetPubKey().(multisig.PubKeyMultisigThreshold)
		multisigSig := multisig.NewMultisig(len(multisigPub.PubKeys))
		clientCtx = clientCtx.InitWithInput(inBuf)
		txFactory := tx.NewFactoryCLI(clientCtx, cmd.Flags())

		if !clientCtx.Offline {
			accnum, seq, err := types.NewAccountRetriever(clientCtx.JSONMarshaler).GetAccountNumberSequence(clientCtx, multisigInfo.GetAddress())
			if err != nil {
				return err
			}

			txFactory = txFactory.WithAccountNumber(accnum).WithSequence(seq)
		}

		feeTx := stdTx.(sdk.FeeTx)
		fee := types.StdFee{
			Amount: feeTx.GetFee(),
			Gas:    feeTx.GetGas(),
		}
		memoTx := stdTx.(sdk.TxWithMemo)

		// read each signature and add it to the multisig if valid
		for i := 2; i < len(args); i++ {
			stdSig, err := readAndUnmarshalStdSignature(clientCtx, args[i])
			if err != nil {
				return err
			}

			signingData := signing.SignerData{
				ChainID:         txFactory.ChainID(),
				AccountNumber:   txFactory.AccountNumber(),
				AccountSequence: txFactory.Sequence(),
			}
			err = signing.VerifySignature(stdSig.PubKey, signingData, stdSig.Data, clientCtx.TxGenerator.SignModeHandler(), stdTx)
			if err != nil {
				return fmt.Errorf("couldn't verify signature")
			}

			if err := multisig.AddSignatureV2(multisigSig, stdSig, multisigPub.PubKeys); err != nil {
				return err
			}
		}

		sigBz, err := types.SignatureDataToAminoSignature(cdc, multisigSig)
		if err != nil {
			return err
		}

		newStdSig := types.StdSignature{Signature: sigBz, PubKey: multisigPub.Bytes()}                   //nolint:staticcheck
		newTx := types.NewStdTx(stdTx.GetMsgs(), fee, []types.StdSignature{newStdSig}, memoTx.GetMemo()) //nolint:staticcheck

		var json []byte

		txBuilder := clientCtx.TxGenerator.NewTxBuilder()
		if err != nil {
			return err
		}

		sigBldr := signingtypes.SignatureV2{
			PubKey: multisigPub,
			Data:   multisigSig,
		}

		err = txBuilder.SetSignatures(sigBldr)

		if err != nil {
			return err
		}
		sigOnly, _ := cmd.Flags().GetBool(flagSigOnly)
		if sigOnly {
			json, err = cdc.MarshalJSON(newTx.Signatures[0])
		} else {
			json, err = cdc.MarshalJSON(newTx)
		}

		if err != nil {
			return err
		}

		outputDoc, _ := cmd.Flags().GetString(flags.FlagOutputDocument)
		if outputDoc == "" {
			cmd.Printf("%s\n", json)
			return
		}

		fp, err := os.OpenFile(outputDoc, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer fp.Close()

		fmt.Fprintf(fp, "%s\n", json)
		return
	}
}

func readAndUnmarshalStdSignature(clientCtx client.Context, filename string) (stdSig signingtypes.SignatureV2, err error) {
	var bytes []byte
	if bytes, err = ioutil.ReadFile(filename); err != nil {
		return
	}
	if err = clientCtx.JSONMarshaler.UnmarshalJSON(bytes, &stdSig); err != nil {
		return
	}
	return
}
