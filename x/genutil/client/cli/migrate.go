package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	tmjson "github.com/tendermint/tendermint/libs/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	v036 "github.com/cosmos/cosmos-sdk/x/genutil/legacy/v036"
	v038 "github.com/cosmos/cosmos-sdk/x/genutil/legacy/v038"
	v039 "github.com/cosmos/cosmos-sdk/x/genutil/legacy/v039"
	v040 "github.com/cosmos/cosmos-sdk/x/genutil/legacy/v040"
	"github.com/cosmos/cosmos-sdk/x/genutil/types"
)

const flagGenesisTime = "genesis-time"

// Allow applications to extend and modify the migration process.
//
// Ref: https://github.com/cosmos/cosmos-sdk/issues/5041
var migrationMap = types.MigrationMap{
	"v0.36": v036.Migrate,
	"v0.38": v038.Migrate, // NOTE: v0.37 and v0.38 are genesis compatible
	"v0.39": v039.Migrate,
	"v0.40": v040.Migrate,
}

// GetMigrationCallback returns a MigrationCallback for a given version.
func GetMigrationCallback(version string) types.MigrationCallback {
	return migrationMap[version]
}

// GetMigrationVersions get all migration version in a sorted slice.
func GetMigrationVersions() []string {
	versions := make([]string, len(migrationMap))

	var i int

	for version := range migrationMap {
		versions[i] = version
		i++
	}

	sort.Strings(versions)

	return versions
}

// MigrateGenesisCmd returns a command to execute genesis state migration.
func MigrateGenesisCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate [target-version] [genesis-file]",
		Short: "Migrate genesis to a specified target version",
		Long: fmt.Sprintf(`Migrate the source genesis into the target version and print to STDOUT.

Example:
$ %s migrate v0.36 /path/to/genesis.json --chain-id=cosmoshub-3 --genesis-time=2019-04-22T17:00:00Z
`, version.AppName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			var err error

			target := args[0]
			importGenesis := args[1]

			jsonBlob, err := ioutil.ReadFile(genDocFile)
			if err != nil {
				return nil, fmt.Errorf("couldn't read GenesisDoc file: %w", err)
			}
			genDoc, err := tmtypes.GenesisDocFromJSON(jsonBlob)
			if err != nil {
				return errors.Wrapf(err, "failed to read genesis document from file %s", importGenesis)
			}

			var initialState types.AppMap
			if err := json.Unmarshal(genDoc.AppState, &initialState); err != nil {
				return errors.Wrap(err, "failed to JSON unmarshal initial genesis state")
			}

			migrationFunc := GetMigrationCallback(target)
			if migrationFunc == nil {
				return fmt.Errorf("unknown migration function for version: %s", target)
			}

			// TODO: handler error from migrationFunc call
			newGenState := migrationFunc(initialState, clientCtx)

			genDoc.AppState, err = json.Marshal(newGenState)
			if err != nil {
				return errors.Wrap(err, "failed to JSON marshal migrated genesis state")
			}

			genesisTime, _ := cmd.Flags().GetString(flagGenesisTime)
			if genesisTime != "" {
				var t time.Time

				err := t.UnmarshalText([]byte(genesisTime))
				if err != nil {
					return errors.Wrap(err, "failed to unmarshal genesis time")
				}

				genDoc.GenesisTime = t
			}

			chainID, _ := cmd.Flags().GetString(flags.FlagChainID)
			if chainID != "" {
				genDoc.ChainID = chainID
			}

			bz, err := tmjson.Marshal(genDoc)
			if err != nil {
				return errors.Wrap(err, "failed to marshal genesis doc")
			}

			sortedBz, err := sdk.SortJSON(bz)
			if err != nil {
				return errors.Wrap(err, "failed to sort JSON genesis doc")
			}

			fmt.Println(string(sortedBz))
			return nil
		},
	}

	cmd.Flags().String(flagGenesisTime, "", "override genesis_time with this flag")
	cmd.Flags().String(flags.FlagChainID, "", "override chain_id with this flag")

	return cmd
}

// SanitizeTendermintGenesis makes sure a later version of Tendermint can parse
// a JSON blob exported by a previous version of Tendermint.
func SanitizeTendermintGenesis(jsonBlob []byte) ([]byte, error) {
	var jsonObj map[string]interface{}
	err := tmjson.Unmarshal(jsonBlob, &jsonObj)
	if err != nil {
		return nil, err
	}

	consensusParams, ok = jsonObj["consensus_params"]["evidence"]
	if !ok {
		return nil, fmt.Errorf("exported json does not contain consensus_params.evidence field")
	}
}
