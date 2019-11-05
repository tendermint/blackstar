package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/evidence/internal/types"

	"github.com/spf13/cobra"
)

// GetTxCmd returns a CLI command that has all the native evidence module tx
// commands mounted. In addition, it mounts all childCmds, implemented by outside
// modules, under a sub-command. This allows external modules to implement custom
// Evidence types and Handlers while having the ability to create and sign txs
// containing them all from a single root command.
func GetTxCmd(storeKey string, cdc *codec.Codec, childCmds []*cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Evidence transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	submitEvidenceCmd := SubmitEvidenceCmd(cdc)
	for _, childCmd := range childCmds {
		submitEvidenceCmd.AddCommand(client.PostCommands(childCmd)[0])
	}

	// TODO: Add tx commands.

	return cmd
}

// SubmitEvidenceCMD returns the top-level evidence submission command handler.
// All concrete evidence submission child command handlers should be registered
// under this command.
func SubmitEvidenceCMD(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit",
		Short: "Submit arbitrary evidence of misbehavior",
	}

	return cmd
}
