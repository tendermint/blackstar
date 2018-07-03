package mock

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	abci "github.com/tendermint/abci/types"
	crypto "github.com/tendermint/go-crypto"
)

// Produce a random transaction, and check the state transition. Return a descriptive message "action" about
// what this random tx actually did, for ease of debugging.
type TestAndRunTx func(t *testing.T, r *rand.Rand, app *App, ctx sdk.Context, privKeys []crypto.PrivKey, log string) (action string, err sdk.Error)

// Perform the random setup the module needs.
type RandSetup func(r *rand.Rand, privKeys []crypto.PrivKey)

// Assert invariants for the module. Print out the log when failing
type AssertInvariants func(t *testing.T, app *App, log string)

// Invariant for Auth module. Placed here to avoid ciruclar dependency
func AuthInvariant(t *testing.T, app *App, log string) {
	// This is a slow check, so only run it 1 out of every 5 blocks
	if app.LastBlockHeight()%5 != 1 {
		return
	}
	ctx := app.BaseApp.NewContext(false, abci.Header{})
	totalCoins := sdk.Coins{}
	chkAccount := func(acc auth.Account) bool {
		coins := acc.GetCoins()
		totalCoins = totalCoins.Plus(coins)
		for _, coin := range coins {
			require.True(t, coin.Amount > 0, log)
		}
		return false
	}
	app.AccountMapper.IterateAccounts(ctx, chkAccount)
	require.Equal(t, app.TotalCoinsSupply, totalCoins, log)
}
