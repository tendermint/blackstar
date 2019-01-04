package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/distribution/types"
)

// initialize rewards for a new validator
func (k Keeper) initializeValidator(ctx sdk.Context, val sdk.Validator) {
	// set initial historical rewards (period 0)
	k.setValidatorHistoricalRewards(ctx, val.GetOperator(), 0, types.ValidatorHistoricalRewards{})

	// set current rewards (starting at period 1)
	k.setValidatorCurrentRewards(ctx, val.GetOperator(), types.ValidatorCurrentRewards{
		Rewards: sdk.DecCoins{},
		Period:  1,
	})

	// set accumulated commission
	k.setValidatorAccumulatedCommission(ctx, val.GetOperator(), types.ValidatorAccumulatedCommission{})
}

// increment validator period, returning the period just ended
func (k Keeper) incrementValidatorPeriod(ctx sdk.Context, val sdk.Validator) uint64 {
	// Fetch current rewards
	rewards := k.GetValidatorCurrentRewards(ctx, val.GetOperator())

	// Calculate current ratio
	var current sdk.DecCoins
	if val.GetPower().IsZero() {
		// this can happen after redelegations are slashed
		// TODO add to the community pool?
		current = sdk.DecCoins{}
	} else {
		current = rewards.Rewards.QuoDec(sdk.NewDecFromInt(val.GetPower()))
	}

	// Fetch historical rewards for last period
	historical := k.GetValidatorHistoricalRewards(ctx, val.GetOperator(), rewards.Period-1)

	// Set new historical rewards
	k.setValidatorHistoricalRewards(ctx, val.GetOperator(), rewards.Period, historical.Plus(current))

	// Set current rewards, incrementing period by 1
	newPeriod := rewards.Period + 1
	k.setValidatorCurrentRewards(ctx, val.GetOperator(), types.ValidatorCurrentRewards{
		Rewards: sdk.DecCoins{},
		Period:  newPeriod,
	})

	return rewards.Period
}

// initialize starting info for a new delegation
func (k Keeper) initializeDelegation(ctx sdk.Context, val sdk.ValAddress, del sdk.AccAddress) {
	// period has already been incremented
	period := k.GetValidatorCurrentRewards(ctx, val).Period
	validator := k.stakeKeeper.Validator(ctx, val)
	delegation := k.stakeKeeper.Delegation(ctx, del, val)
	// calculate delegation stake in tokens
	// TODO need to make sure this truncates instead of rounding
	stake := delegation.GetShares().Mul(validator.GetDelegatorShareExRate())
	k.setDelegatorStartingInfo(ctx, val, del, types.DelegatorStartingInfo{
		PreviousPeriod: period - 1,
		Stake:          stake,
		Height:         uint64(ctx.BlockHeight()),
	})
}

// calculate the rewards accrued by a delegation between two periods
func (k Keeper) calculateDelegationRewardsBetween(ctx sdk.Context, val sdk.Validator, startingPeriod uint64, endingPeriod uint64, stake sdk.Dec) (rewards sdk.DecCoins) {
	// sanity check
	if startingPeriod > endingPeriod {
		panic("startingPeriod cannot be greater than endingPeriod")
	}

	// return stake * (ending - starting)
	starting := k.GetValidatorHistoricalRewards(ctx, val.GetOperator(), startingPeriod)
	ending := k.GetValidatorHistoricalRewards(ctx, val.GetOperator(), endingPeriod)
	difference := ending.Minus(starting)
	rewards = difference.MulDec(stake)
	return
}

// calculate the total rewards accrued by a delegation
func (k Keeper) calculateDelegationRewards(ctx sdk.Context, val sdk.Validator, del sdk.Delegation, endingPeriod uint64) (rewards sdk.DecCoins) {
	// fetch starting info for delegation
	startingInfo := k.GetDelegatorStartingInfo(ctx, del.GetValidatorAddr(), del.GetDelegatorAddr())
	startingPeriod := startingInfo.PreviousPeriod
	stake := startingInfo.Stake

	// iterate through slashes and withdraw with calculated stake for sub-intervals
	startingHeight := startingInfo.Height + 1
	endingHeight := uint64(ctx.BlockHeight())
	// if we're still in the same block, no slashes can yet have happened
	if endingHeight > startingHeight {
		k.IterateValidatorSlashEvents(ctx, del.GetValidatorAddr(), startingHeight, endingHeight, func(height uint64, event types.ValidatorSlashEvent) (stop bool) {
			stake = stake.Mul(sdk.OneDec().Sub(event.Fraction))
			endingPeriod := event.ValidatorPeriod
			rewards = rewards.Plus(k.calculateDelegationRewardsBetween(ctx, val, startingPeriod, endingPeriod, stake))
			startingPeriod = endingPeriod
			return false
		})
	}

	// calculate rewards for final period
	rewards = rewards.Plus(k.calculateDelegationRewardsBetween(ctx, val, startingPeriod, endingPeriod, stake))
	return
}

func (k Keeper) withdrawDelegationRewards(ctx sdk.Context, val sdk.Validator, del sdk.Delegation) sdk.Error {

	// end current period and calculate rewards
	endingPeriod := k.incrementValidatorPeriod(ctx, val)
	rewards := k.calculateDelegationRewards(ctx, val, del, endingPeriod)

	// truncate coins, return remainder to community pool
	coins, remainder := rewards.TruncateDecimal()
	outstanding := k.GetOutstandingRewards(ctx)
	k.SetOutstandingRewards(ctx, outstanding.Minus(rewards))
	feePool := k.GetFeePool(ctx)
	feePool.CommunityPool = feePool.CommunityPool.Plus(remainder)
	k.SetFeePool(ctx, feePool)

	// add coins to user account
	withdrawAddr := k.GetDelegatorWithdrawAddr(ctx, del.GetDelegatorAddr())
	if _, _, err := k.bankKeeper.AddCoins(ctx, withdrawAddr, coins); err != nil {
		return err
	}

	return nil
}

func (k Keeper) updateValidatorSlashFraction(ctx sdk.Context, valAddr sdk.ValAddress, fraction sdk.Dec) {
	height := uint64(ctx.BlockHeight())
	currentFraction := sdk.ZeroDec()
	currentPeriod := k.GetValidatorCurrentRewards(ctx, valAddr).Period
	current, found := k.GetValidatorSlashEvent(ctx, valAddr, height)
	if found {
		currentFraction = current.Fraction
	}
	updatedFraction := sdk.OneDec().Sub(sdk.OneDec().Sub(currentFraction).Mul(sdk.OneDec().Sub(fraction)))
	k.setValidatorSlashEvent(ctx, valAddr, height, types.ValidatorSlashEvent{
		ValidatorPeriod: currentPeriod,
		Fraction:        updatedFraction,
	})
}
