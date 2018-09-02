package tx

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"

	"github.com/tendermint/tendermint/libs/common"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	abci "github.com/tendermint/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/auth"
	tmliteErr "github.com/tendermint/tendermint/lite/errors"
	tmliteProxy "github.com/tendermint/tendermint/lite/proxy"
)

// QueryTxCmd implements the default command for a tx query.
func QueryTxCmd(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tx [hash]",
		Short: "Matches this txhash over all committed blocks",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// find the key to look up the account
			hashHexStr := args[0]
			distrustNode := viper.GetBool(client.FlagDistrustNode)

			cliCtx := context.NewCLIContext().WithCodec(cdc)

			output, err := queryTx(cdc, cliCtx, hashHexStr, distrustNode)
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	cmd.Flags().StringP(client.FlagNode, "n", "tcp://localhost:26657", "Node to connect to")
	cmd.Flags().String(client.FlagChainID, "", "The chain ID to connect to")
	cmd.Flags().Bool(client.FlagDistrustNode, true, "Verify proofs for query responses if true")
	return cmd
}

func queryTx(cdc *wire.Codec, cliCtx context.CLIContext, hashHexStr string, distrustNode bool) ([]byte, error) {
	hash, err := hex.DecodeString(hashHexStr)
	if err != nil {
		return nil, err
	}

	node, err := cliCtx.GetNode()
	if err != nil {
		return nil, err
	}

	res, err := node.Tx(hash, distrustNode)
	if err != nil {
		return nil, err
	}

	info, err := formatTxResult(cdc, res)
	if err != nil {
		return nil, err
	}

	if distrustNode {
		check, err := tmliteProxy.GetCertifiedCommit(info.Height, node, cliCtx.Certifier)
		if tmliteErr.IsCommitNotFoundErr(err) {
			return nil, context.ErrVerifyCommit(info.Height)
		} else if err != nil {
			return nil, err
		}

		err = res.Proof.Validate(check.Header.DataHash)
		if err != nil {
			return nil, err
		}
	}

	return wire.MarshalJSONIndent(cdc, info)
}

func formatTxResult(cdc *wire.Codec, res *ctypes.ResultTx) (Info, error) {
	// TODO: verify the proof if requested
	tx, err := parseTx(cdc, res.Tx)
	if err != nil {
		return Info{}, err
	}

	return Info{
		Hash:   res.Hash,
		Height: res.Height,
		Tx:     tx,
		Result: res.TxResult,
	}, nil
}

// Info is used to prepare info to display
type Info struct {
	Hash   common.HexBytes        `json:"hash"`
	Height int64                  `json:"height"`
	Tx     sdk.Tx                 `json:"tx"`
	Result abci.ResponseDeliverTx `json:"result"`
}

func parseTx(cdc *wire.Codec, txBytes []byte) (sdk.Tx, error) {
	var tx auth.StdTx

	err := cdc.UnmarshalBinary(txBytes, &tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// REST

// transaction query REST handler
func QueryTxRequestHandlerFn(cdc *wire.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		hashHexStr := vars["hash"]
		distrustNode, err := strconv.ParseBool(r.FormValue("distrust_node"))
		// distrustNode defaults to true
		if err != nil {
			distrustNode = true
		}

		output, err := queryTx(cdc, cliCtx, hashHexStr, distrustNode)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(output)
	}
}
