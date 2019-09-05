package keys

import (
	"bufio"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func migrateKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate key information from the lagacy key database to the OS secret store, or encrypted file store as a fall-back and save it",
		Long: `This command migrates key information from the legacy secret store to the OS secret store. The command asks for every passphrase. 
		If passphrase is incorrect, it skips the key. 

		Previous versions of Cosmos SDK used a custom secret store. On version xxx, Cosmos SDK CLI was updated to use a library Keyring (https://github.com/99designs/keyring) to 
		preferentially store secrets in the secret manager of many Operating Systems. This is intended to provide stronger security guarantees than the 
		custom secret store is provided. 
`,
		Args: cobra.ExactArgs(0),
		RunE: runMigrateCmd,
	}

	return cmd
}

func runMigrateCmd(cmd *cobra.Command, args []string) error {

	var legacykb keys.Keybase
	var keyringkb keys.Keybase
	var err error

	//instantiating variables
	rootDir := viper.GetString(flags.FlagHome)
	legacykb, err = NewKeyBaseFromHomeFlag()
	if err != nil {
		return err
	}

	buf := bufio.NewReader(cmd.InOrStdin())

	keyringkb = keys.NewKeybaseKeyring(types.GetConfig().GetKeyringServiceName(), rootDir, buf, false)

	legacyKeyList, err := legacykb.List()
	for _, key := range legacyKeyList {

		legKeyInfo, err := legacykb.Export(key.GetName())
		if err != nil {
			return err
		}

		switch key.GetType() {
		case keys.TypeLocal:
			cmd.PrintErrf("Migrating %s \n", key.GetName())
			decryptPassword, err := input.GetPassword("Enter passphrase to decrypt your key:", buf)
			if err != nil {
				return err
			}

			privData, err := legacykb.ExportPrivKey(key.GetName(), decryptPassword, "abc")
			if err != nil {
				return err
			}

			err = keyringkb.ImportPrivKey(key.GetName(), privData, "abc")
			if err != nil {
				return err
			}

		case keys.TypeOffline, keys.TypeMulti, keys.TypeLedger:
			err = keyringkb.Import(key.GetName(), legKeyInfo)
			if err != nil {
				return err
			}

		}
	}

	return err
}
