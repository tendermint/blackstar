package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	commitmenttypes "github.com/cosmos/cosmos-sdk/x/ibc/23-commitment/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/common"
)

func TestValidateGenesis(t *testing.T) {

	testCases := []struct {
		name     string
		genState GenesisState
		expPass  bool
	}{
		{
			name:     "default",
			genState: DefaultGenesisState(),
			expPass:  true,
		},
		{
			name: "valid genesis",
			genState: NewGenesisState(
				[]ConnectionEnd{
					NewConnectionEnd(common.INIT, connectionID, clientID, Counterparty{clientID2, connectionID2, commitmenttypes.NewMerklePrefix([]byte("prefix"))}, []string{"1.0.0"}),
				},
				[]ConnectionPaths{
					{clientID, []string{common.ConnectionPath(connectionID)}},
				},
			),
			expPass: true,
		},
		{
			name: "invalid connection",
			genState: NewGenesisState(
				[]ConnectionEnd{
					NewConnectionEnd(common.INIT, connectionID, "CLIENTIDONE", Counterparty{clientID, connectionID, commitmenttypes.NewMerklePrefix([]byte("prefix"))}, []string{"1.0.0"}),
				},
				[]ConnectionPaths{
					{clientID, []string{common.ConnectionPath(connectionID)}},
				},
			),
			expPass: false,
		},
		{
			name: "invalid client id",
			genState: NewGenesisState(
				[]ConnectionEnd{
					NewConnectionEnd(common.INIT, connectionID, clientID, Counterparty{clientID2, connectionID2, commitmenttypes.NewMerklePrefix([]byte("prefix"))}, []string{"1.0.0"}),
				},
				[]ConnectionPaths{
					{"CLIENTIDONE", []string{common.ConnectionPath(connectionID)}},
				},
			),
			expPass: false,
		},
		{
			name: "invalid path",
			genState: NewGenesisState(
				[]ConnectionEnd{
					NewConnectionEnd(common.INIT, connectionID, clientID, Counterparty{clientID2, connectionID2, commitmenttypes.NewMerklePrefix([]byte("prefix"))}, []string{"1.0.0"}),
				},
				[]ConnectionPaths{
					{clientID, []string{connectionID}},
				},
			),
			expPass: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		err := tc.genState.Validate()
		if tc.expPass {
			require.NoError(t, err, tc.name)
		} else {
			require.Error(t, err, tc.name)
		}
	}
}
