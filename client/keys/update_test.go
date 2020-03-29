package keys

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_updateKeyCommand(t *testing.T) {
	require.NotNil(t, UpdateKeyCommand())
	// No flags  or defaults to validate
}

// func Test_runUpdateCmd(t *testing.T) {
// 	fakeKeyName1 := "runUpdateCmd_Key1"
// 	fakeKeyName2 := "runUpdateCmd_Key2"
//
// 	cmd := UpdateKeyCommand()
//
// 	// fails because it requests a password
// 	err := runUpdateCmd(cmd, []string{fakeKeyName1})
//
// 	require.EqualError(t, err, "EOF")
//
// 	// try again
// 	mockIn, _, _ := tests.ApplyMockIO(cmd)
// 	mockIn.Reset("pass1234\n")
// 	err = runUpdateCmd(cmd, []string{fakeKeyName1})
// 	require.Error(t, err)
//
// 	// Prepare a key base
// 	// Now add a temporary keybase
// 	kbHome, cleanUp1 := tests.NewTestCaseDir(t)
// 	t.Cleanup(cleanUp1)
// 	viper.Set(flags.FlagHome, kbHome)
//
// 	kb, err := NewKeyBaseFromDir(kbHome)
// 	require.NoError(t, err)
// 	_, err = kb.CreateAccount(fakeKeyName1, tests.TestMnemonic, "", "", "0", keyring.Secp256k1)
// 	require.NoError(t, err)
// 	_, err = kb.CreateAccount(fakeKeyName2, tests.TestMnemonic, "", "", "1", keyring.Secp256k1)
// 	require.NoError(t, err)
//
// 	// Try again now that we have keys
// 	// Incorrect key type
// 	mockIn.Reset("pass1234\nNew1234\nNew1234")
// 	err = runUpdateCmd(cmd, []string{fakeKeyName1})
// 	require.EqualError(t, err, "locally stored key required. Received: keyring.offlineInfo")
//
// 	// TODO: Check for other type types?
// }
