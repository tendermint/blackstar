package app

import (
	"encoding/json"
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	stake "github.com/cosmos/cosmos-sdk/x/stake"
)

// export the state of gaia for a genesis file
func (app *GaiaApp) ExportAppStateAndValidators(forZeroHeight bool) (
	appState json.RawMessage, validators []tmtypes.GenesisValidator, err error) {

	// as if they could withdraw from the start of the next block
	ctx := app.NewContext(true, abci.Header{Height: app.LastBlockHeight()})

	if forZeroHeight {
		app.prepForZeroHeightGenesis(ctx)
	}

	// iterate to get the accounts
	accounts := []GenesisAccount{}
	appendAccount := func(acc auth.Account) (stop bool) {
		account := NewGenesisAccountI(acc)
		accounts = append(accounts, account)
		return false
	}
	app.accountKeeper.IterateAccounts(ctx, appendAccount)

	genState := NewGenesisState(
		accounts,
		auth.ExportGenesis(ctx, app.accountKeeper, app.feeCollectionKeeper),
		stake.ExportGenesis(ctx, app.stakeKeeper),
		mint.ExportGenesis(ctx, app.mintKeeper),
		distr.ExportGenesis(ctx, app.distrKeeper),
		gov.ExportGenesis(ctx, app.govKeeper),
		slashing.ExportGenesis(ctx, app.slashingKeeper),
	)
	appState, err = codec.MarshalJSONIndent(app.cdc, genState)
	if err != nil {
		return nil, nil, err
	}
	validators = stake.WriteValidators(ctx, app.stakeKeeper)
	return appState, validators, nil
}

// prepare for fresh start at zero height
func (app *GaiaApp) prepForZeroHeightGenesis(ctx sdk.Context) {

	/* Just to be safe, assert the invariants on current state. */
	app.assertRuntimeInvariantsOnContext(ctx)

	/* Handle fee distribution state. */

	// withdraw all delegator & validator rewards
	vdiIter := func(_ int64, valInfo distr.ValidatorDistInfo) (stop bool) {
		err := app.distrKeeper.WithdrawValidatorRewardsAll(ctx, valInfo.OperatorAddr)
		if err != nil {
			panic(err)
		}
		return false
	}
	app.distrKeeper.IterateValidatorDistInfos(ctx, vdiIter)

	ddiIter := func(_ int64, distInfo distr.DelegationDistInfo) (stop bool) {
		err := app.distrKeeper.WithdrawDelegationReward(
			ctx, distInfo.DelegatorAddr, distInfo.ValOperatorAddr)
		if err != nil {
			panic(err)
		}
		return false
	}
	app.distrKeeper.IterateDelegationDistInfos(ctx, ddiIter)

	app.assertRuntimeInvariantsOnContext(ctx)

	// set distribution info withdrawal heights to 0
	app.distrKeeper.IterateDelegationDistInfos(ctx, func(_ int64, delInfo distr.DelegationDistInfo) (stop bool) {
		delInfo.DelPoolWithdrawalHeight = 0
		app.distrKeeper.SetDelegationDistInfo(ctx, delInfo)
		return false
	})
	app.distrKeeper.IterateValidatorDistInfos(ctx, func(_ int64, valInfo distr.ValidatorDistInfo) (stop bool) {
		valInfo.FeePoolWithdrawalHeight = 0
		valInfo.DelAccum.UpdateHeight = 0
		app.distrKeeper.SetValidatorDistInfo(ctx, valInfo)
		return false
	})

	// assert that the fee pool is empty
	feePool := app.distrKeeper.GetFeePool(ctx)
	if !feePool.TotalValAccum.Accum.IsZero() {
		panic("unexpected leftover validator accum")
	}
	bondDenom := app.stakeKeeper.GetParams(ctx).BondDenom
	if !feePool.ValPool.AmountOf(bondDenom).IsZero() {
		panic(fmt.Sprintf("unexpected leftover validator pool coins: %v",
			feePool.ValPool.AmountOf(bondDenom).String()))
	}

	// reset fee pool height, save fee pool
	feePool.TotalValAccum = distr.NewTotalAccum(0)
	app.distrKeeper.SetFeePool(ctx, feePool)

	/* Handle stake state. */

	// iterate through redelegations, reset creation height
	app.stakeKeeper.IterateRedelegations(ctx, func(_ int64, red stake.Redelegation) (stop bool) {
		red.CreationHeight = 0
		app.stakeKeeper.SetRedelegation(ctx, red)
		return false
	})

	// iterate through unbonding delegations, reset creation height
	app.stakeKeeper.IterateUnbondingDelegations(ctx, func(_ int64, ubd stake.UnbondingDelegation) (stop bool) {
		ubd.CreationHeight = 0
		app.stakeKeeper.SetUnbondingDelegation(ctx, ubd)
		return false
	})

	// Iterate through validators by power descending, reset bond heights, and
	// update bond intra-tx counters.
	store := ctx.KVStore(app.keyStake)
	iter := sdk.KVStoreReversePrefixIterator(store, stake.ValidatorsKey)
	counter := int16(0)

	var valConsAddrs []sdk.ConsAddress
	for ; iter.Valid(); iter.Next() {
		addr := sdk.ValAddress(iter.Key()[1:])
		validator, found := app.stakeKeeper.GetValidator(ctx, addr)
		if !found {
			panic("expected validator, not found")
		}

		validator.BondHeight = 0
		validator.UnbondingHeight = 0
		valConsAddrs = append(valConsAddrs, validator.ConsAddress())

		app.stakeKeeper.SetValidator(ctx, validator)
		counter++
	}

	iter.Close()

	/* Handle slashing state. */

	// remove all existing slashing periods and recreate one for each validator
	app.slashingKeeper.DeleteValidatorSlashingPeriods(ctx)

	for _, valConsAddr := range valConsAddrs {
		sp := slashing.ValidatorSlashingPeriod{
			ValidatorAddr: valConsAddr,
			StartHeight:   0,
			EndHeight:     0,
			SlashedSoFar:  sdk.ZeroDec(),
		}
		app.slashingKeeper.SetValidatorSlashingPeriod(ctx, sp)
	}

	// reset slashing periods
	app.slashingKeeper.IterateValidatorSlashingPeriods(
		ctx,
		func(sp slashing.ValidatorSlashingPeriod) (stop bool) {
			sp.StartHeight = 0
			return false
		},
	)

	// reset start height on signing infos
	app.slashingKeeper.IterateValidatorSigningInfos(
		ctx,
		func(addr sdk.ConsAddress, info slashing.ValidatorSigningInfo) (stop bool) {
			info.StartHeight = 0
			app.slashingKeeper.SetValidatorSigningInfo(ctx, addr, info)
			return false
		},
	)
}
