package keys

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/crypto/hd"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func Test_runExportCmd(t *testing.T) {
	cmd := ExportKeyCommand()
	cmd.Flags().AddFlagSet(Commands().PersistentFlags())
	mockIn, _, _ := tests.ApplyMockIO(cmd)

	// Now add a temporary keybase
	kbHome, cleanUp := tests.NewTestCaseDir(t)
	t.Cleanup(cleanUp)

	// create a key
	kb, err := keyring.New(sdk.KeyringServiceName(), keyring.BackendTest, kbHome, mockIn)
	require.NoError(t, err)
	t.Cleanup(func() {
		kb.Delete("keyname1") // nolint:errcheck
	})

	path := sdk.GetConfig().GetFullFundraiserPath()
	_, err = kb.NewAccount("keyname1", tests.TestMnemonic, "", path, hd.Secp256k1)
	require.NoError(t, err)

	// Now enter password
	cmd.SetArgs([]string{"keyname1", fmt.Sprintf("--%s=%s", flags.FlagHome, kbHome)})
	mockIn.Reset("123456789\n123456789\n")
	require.NoError(t, cmd.Execute())
}
