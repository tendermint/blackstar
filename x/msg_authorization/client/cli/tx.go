package cli

import (
	"bufio"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/msg_authorization/internal/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	AuthorizationTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Authorization transactions subcommands",
		Long:                       "Authorize and revoke access to execute transactions on behalf of your address",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	AuthorizationTxCmd.AddCommand(client.PostCommands(
		GetCmdGrantCapability(cdc),
		GetCmdRevokeCapability(cdc),
	)...)

	return AuthorizationTxCmd
}

func GetCmdGrantCapability(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grant",
		Short: "Grant authorization to an address",
		Long:  "Grant authorization to an address to execute a transaction on your behalf",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			granter := cliCtx.FromAddress
			grantee, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			var capability types.Capability
			err = cdc.UnmarshalJSON([]byte(args[1]), &capability)
			if err != nil {
				return err
			}
			expirationString := viper.GetString(FlagExpiration)
			expiration, err := time.Parse(time.RFC3339, expirationString)
			if err != nil {
				return err
			}

			msg := types.NewMsgGrantAuthorization(granter, grantee, capability, expiration)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})

		},
	}
	cmd.Flags().String(FlagExpiration, "9999-12-31 23:59:59.52Z", "The time upto which the authorization is active for the user")

	return cmd
}

func GetCmdRevokeCapability(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke",
		Short: "revoke authorization",
		Long:  "revoke authorization from an address for a transaction",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			granter := cliCtx.FromAddress
			grantee, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			var msgAuthorized sdk.Msg
			err = cdc.UnmarshalJSON([]byte(args[1]), &msgAuthorized)
			if err != nil {
				return err
			}

			msg := types.NewMsgRevokeAuthorization(granter, grantee, msgAuthorized)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
	return cmd
}
