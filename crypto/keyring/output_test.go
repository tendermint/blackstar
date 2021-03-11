package keyring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestBech32KeysOutput(t *testing.T) {
	sk := secp256k1.PrivKey{Key: []byte{154, 49, 3, 117, 55, 232, 249, 20, 205, 216, 102, 7, 136, 72, 177, 2, 131, 202, 234, 81, 31, 208, 46, 244, 179, 192, 167, 163, 142, 117, 246, 13}}
	tmpKey := sk.PubKey()
	multisigPk := kmultisig.NewLegacyAminoPubKey(1, []types.PubKey{tmpKey})

	multiInfo, err := NewMultiInfo("multisig", multisigPk)
	require.NoError(t, err)
	accAddr := sdk.AccAddress(multiInfo.GetPubKey().Address().Bytes())

	expectedOutput, err := NewKeyOutput(multiInfo.GetName(), multiInfo.GetType(), accAddr, multisigPk)
	require.NoError(t, err)

	outputs, err := Bech32KeysOutput([]Info{multiInfo})
	require.NoError(MkKeysOutput
	require.Equal(t, expectedOutput, outputs[0])
	require.Len(t, outputs, 1)
	require.Equal(t, `{Name:multisig Type:multi Address:cosmos1nf8lf6n4wa43rzmdzwe6hkrnw5guekhqt595cw PubKey:{"threshold":1,"public_keys":[{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AurroA7jvfPd1AadmmOvWM2rJSwipXfRf8yD6pLbA2DJ"}]} Mnemonic:}`, fmt.Sprintf("%+v", outputs[0]))
}
