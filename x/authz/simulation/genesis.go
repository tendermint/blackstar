package simulation

import (
	"math/rand"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// genAuthorizationGrant returns a slice of authorization grants.
func genAuthorizationGrant(r *rand.Rand, accounts []simtypes.Account) []authz.GrantAuthorization {
	authorizations := make([]authz.GrantAuthorization, len(accounts)-1)
	for i := 0; i < len(accounts)-1; i++ {
		granter := accounts[i]
		grantee := accounts[i+1]
		authorizations[i] = authz.GrantAuthorization{
			Granter:       granter.Address.String(),
			Grantee:       grantee.Address.String(),
			Authorization: generateRandomGrant(r),
		}
	}

	return authorizations
}

func generateRandomGrant(r *rand.Rand) *codectypes.Any {
	authorizations := make([]*codectypes.Any, 2)
	authorizations[0] = newAnyAuthorization(banktypes.NewSendAuthorization(sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000)))))
	authorizations[1] = newAnyAuthorization(authz.NewGenericAuthorization(sdk.MsgTypeURL(&govtypes.MsgSubmitProposal{})))

	return authorizations[r.Intn(len(authorizations))]
}

func newAnyAuthorization(grant authz.Authorization) *codectypes.Any {
	any, err := codectypes.NewAnyWithValue(grant)
	if err != nil {
		panic(err)
	}

	return any
}

// RandomizedGenState generates a random GenesisState for authz.
func RandomizedGenState(simState *module.SimulationState) {
	var grants []authz.GrantAuthorization
	simState.AppParams.GetOrGenerate(
		simState.Cdc, "authz", &grants, simState.Rand,
		func(r *rand.Rand) { grants = genAuthorizationGrant(r, simState.Accounts) },
	)

	authzGrantsGenesis := authz.NewGenesisState(grants)
	authzGrantsGenesis = authz.NewGenesisState([]authz.GrantAuthorization{})

	simState.GenState[authz.ModuleName] = simState.Cdc.MustMarshalJSON(authzGrantsGenesis)
}
