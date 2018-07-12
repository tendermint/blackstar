package keeper

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/stake/types"
)

//changing the int in NewSource will allow you to test different, deterministic, sets of operations
var r = rand.New(rand.NewSource(6595))

func TestGetInflation(t *testing.T) {
	ctx, _, keeper := CreateTestInput(t, false, 0)
	pool := keeper.GetPool(ctx)
	params := keeper.GetParams(ctx)
	hrsPerYrRat := sdk.NewRat(hrsPerYr)

	// Governing Mechanism:
	//    BondedRatio = BondedTokens / TotalSupply
	//    inflationRateChangePerYear = (1- BondedRatio/ GoalBonded) * MaxInflationRateChange

	tests := []struct {
		name string
		setBondedTokens, setLooseTokens,
		setInflation, expectedChange sdk.Rat
	}{
		// with 0% bonded atom supply the inflation should increase by InflationRateChange
		{"test 1", sdk.ZeroRat(), sdk.ZeroRat(), sdk.NewRat(7, 100), params.InflationRateChange.Quo(hrsPerYrRat).Round(precision)},

		// 100% bonded, starting at 20% inflation and being reduced
		// (1 - (1/0.67))*(0.13/8667)
		{"test 2", sdk.OneRat(), sdk.ZeroRat(), sdk.NewRat(20, 100),
			sdk.OneRat().Sub(sdk.OneRat().Quo(params.GoalBonded)).Mul(params.InflationRateChange).Quo(hrsPerYrRat).Round(precision)},

		// 50% bonded, starting at 10% inflation and being increased
		{"test 3", sdk.OneRat(), sdk.OneRat(), sdk.NewRat(10, 100),
			sdk.OneRat().Sub(sdk.NewRat(1, 2).Quo(params.GoalBonded)).Mul(params.InflationRateChange).Quo(hrsPerYrRat).Round(precision)},

		// test 7% minimum stop (testing with 100% bonded)
		{"test 4", sdk.OneRat(), sdk.ZeroRat(), sdk.NewRat(7, 100), sdk.ZeroRat()},
		{"test 5", sdk.OneRat(), sdk.ZeroRat(), sdk.NewRat(70001, 1000000), sdk.NewRat(-1, 1000000).Round(precision)},

		// test 20% maximum stop (testing with 0% bonded)
		{"test 6", sdk.ZeroRat(), sdk.ZeroRat(), sdk.NewRat(20, 100), sdk.ZeroRat()},
		{"test 7", sdk.ZeroRat(), sdk.ZeroRat(), sdk.NewRat(199999, 1000000), sdk.NewRat(1, 1000000).Round(precision)},

		// perfect balance shouldn't change inflation
		{"test 8", sdk.NewRat(67), sdk.NewRat(33), sdk.NewRat(15, 100), sdk.ZeroRat()},
	}
	for _, tc := range tests {
		pool.BondedTokens, pool.LooseTokens = tc.setBondedTokens, tc.setLooseTokens
		pool.Inflation = tc.setInflation
		keeper.SetPool(ctx, pool)

		inflation := keeper.NextInflation(ctx)
		diffInflation := inflation.Sub(tc.setInflation)

		require.True(t, diffInflation.Equal(tc.expectedChange),
			"Name: %v\nDiff:  %v\nExpected: %v\n", tc.name, diffInflation, tc.expectedChange)
	}
}

// Test that provisions are correctly added to the pool and validators each hour for 1 year
func TestProcessProvisions(t *testing.T) {
	ctx, _, keeper := CreateTestInput(t, false, 0)
	pool := keeper.GetPool(ctx)

	var (
		initialTotalTokens  int64  = 550000000
		initialBondedTokens int64  = 250000000
		cumulativeExpProvs         = sdk.ZeroRat()
		validatorTokens            = []int64{150000000, 100000000, 100000000, 100000000, 100000000}
		bondedValidators    uint16 = 2
	)
	pool.LooseTokens = sdk.NewRat(initialTotalTokens)

	// create some validators some bonded, some unbonded
	_, keeper, pool = setupTestValidators(pool, keeper, ctx, validatorTokens, bondedValidators)
	checkValidatorSetup(t, pool, initialTotalTokens, initialBondedTokens)

	// process the provisions for a year
	for hr := 0; hr < 8766; hr++ {
		pool := keeper.GetPool(ctx)
		_, expProvisions, _ := updateProvisions(t, keeper, pool, ctx, hr)
		cumulativeExpProvs = cumulativeExpProvs.Add(expProvisions)
	}

	//get the pool and do the final value checks from checkFinalPoolValues
	pool = keeper.GetPool(ctx)
	checkFinalPoolValues(t, pool, sdk.NewRat(initialTotalTokens), cumulativeExpProvs)
}

// Tests that the hourly rate of change of inflation will be positive, negative, or zero, depending on bonded ratio and inflation rate
// Cycles through the whole gambit of inflation possibilities, starting at 7% inflation, up to 20%, back down to 7% (it takes ~11.4 years)
func TestHourlyInflationRateOfChange(t *testing.T) {
	ctx, _, keeper := CreateTestInput(t, false, 0)
	pool := keeper.GetPool(ctx)

	var (
		initialTotalTokens  int64  = 550000000
		initialBondedTokens int64  = 150000000
		cumulativeExpProvs         = sdk.ZeroRat()
		validatorTokens            = []int64{150000000, 100000000, 100000000, 100000000, 100000000}
		bondedValidators    uint16 = 1
	)
	pool.LooseTokens = sdk.NewRat(initialTotalTokens)

	// create some validators some bonded, some unbonded
	_, keeper, pool = setupTestValidators(pool, keeper, ctx, validatorTokens, bondedValidators)
	checkValidatorSetup(t, pool, initialTotalTokens, initialBondedTokens)

	// ~11.4 years to go from 7%, up to 20%, back down to 7%
	for hr := 0; hr < 100000; hr++ {
		pool := keeper.GetPool(ctx)
		previousInflation := pool.Inflation
		updatedInflation, expProvisions, pool := updateProvisions(t, keeper, pool, ctx, hr)
		cumulativeExpProvs = cumulativeExpProvs.Add(expProvisions)
		msg := strconv.Itoa(hr)
		checkInflation(t, pool, previousInflation, updatedInflation, msg)
	}

	// Final check that the pool equals initial values + cumulative provisions and adjustments we recorded
	pool = keeper.GetPool(ctx)
	checkFinalPoolValues(t, pool, sdk.NewRat(initialTotalTokens), cumulativeExpProvs)
}

//Test that a large unbonding will significantly lower the bonded ratio
func TestLargeUnbond(t *testing.T) {
	ctx, _, keeper := CreateTestInput(t, false, 0)
	pool := keeper.GetPool(ctx)

	var (
		initialTotalTokens  int64  = 1200000000
		initialBondedTokens int64  = 900000000
		validatorTokens            = []int64{300000000, 100000000, 100000000, 100000000, 100000000, 100000000, 100000000, 100000000, 100000000, 100000000}
		bondedValidators    uint16 = 7
	)
	pool.LooseTokens = sdk.NewRat(initialTotalTokens)

	_, keeper, pool = setupTestValidators(pool, keeper, ctx, validatorTokens, bondedValidators)
	checkValidatorSetup(t, pool, initialTotalTokens, initialBondedTokens)

	pool = keeper.GetPool(ctx)
	validator, found := keeper.GetValidator(ctx, Addrs[0])
	require.True(t, found)

	// initialBondedRatio that we can use to compare to the new values after the unbond
	initialBondedRatio := pool.BondedRatio()

	// validator[0] will be unbonded, bringing us from 75% bonded ratio to ~50% (unbonding 300,000,000)
	pool, validator, _, _ = types.OpBondOrUnbond(r, pool, validator)
	keeper.SetPool(ctx, pool)

	// process provisions after the bonding, to compare the difference in expProvisions and expInflation
	_, expProvisionsAfter, pool := updateProvisions(t, keeper, pool, ctx, 0)

	// unbonded shares should increase
	require.Equal(t, sdk.Unbonded, validator.Status)
	require.True(t, validator.Tokens.GT(sdk.NewRat(300000000, 1)))

	// Ensure that new bonded ratio is less than old bonded ratio , because before they were increasing (i.e. 50% < 75)
	require.True(t, (pool.BondedRatio().LT(initialBondedRatio)))

	// Final check that the pool equals initial values + provisions and adjustments we recorded
	pool = keeper.GetPool(ctx)
	checkFinalPoolValues(t, pool, sdk.NewRat(initialTotalTokens), expProvisionsAfter)
}

//Test that a large bonding will significantly increase the bonded ratio
func TestLargeBond(t *testing.T) {
	ctx, _, keeper := CreateTestInput(t, false, 0)
	pool := keeper.GetPool(ctx)

	var (
		initialTotalTokens  int64  = 1600000000
		initialBondedTokens int64  = 400000000
		validatorTokens            = []int64{400000000, 100000000, 100000000, 100000000, 100000000, 100000000, 100000000, 100000000, 100000000, 400000000}
		bondedValidators    uint16 = 1
	)
	pool.LooseTokens = sdk.NewRat(initialTotalTokens)

	_, keeper, pool = setupTestValidators(pool, keeper, ctx, validatorTokens, bondedValidators)
	checkValidatorSetup(t, pool, initialTotalTokens, initialBondedTokens)

	pool = keeper.GetPool(ctx)
	validator, found := keeper.GetValidator(ctx, Addrs[9])
	require.True(t, found)

	// initialBondedRatio that we can use to compare to the new values after the unbond
	initialBondedRatio := pool.BondedRatio()

	params := types.DefaultParams()
	params.MaxValidators = bondedValidators + 1 //must do this to allow for an extra validator to bond
	keeper.SetParams(ctx, params)

	// validator[9] will be bonded, bringing us from 25% to ~50% (bonding 400,000,000 tokens)
	pool, _, _, _ = types.OpBondOrUnbond(r, pool, validator)
	keeper.SetPool(ctx, pool)

	// process provisions after the bonding, to compare the difference in expProvisions and expInflation
	_, expProvisionsAfter, pool := updateProvisions(t, keeper, pool, ctx, 0)

	// Ensure that new bonded ratio is greater than old bonded ratio (i.e. 50% > 25%)
	require.True(t, (pool.BondedRatio().GT(initialBondedRatio)))
	// Final check that the pool equals initial values + provisions and adjustments we recorded
	pool = keeper.GetPool(ctx)

	checkFinalPoolValues(t, pool, sdk.NewRat(initialTotalTokens), expProvisionsAfter)
}

// Tests that inflation increases or decreases as expected when we do a random operation on 20 different validators
func TestInflationWithRandomOperations(t *testing.T) {
	ctx, _, keeper := CreateTestInput(t, false, 0)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)
	numValidators := 20

	// start off by randomly setting up 20 validators
	pool, validators := types.RandomSetup(r, numValidators)
	require.Equal(t, numValidators, len(validators))

	for i := 0; i < len(validators); i++ {
		keeper.SetValidator(ctx, validators[i])
	}
	keeper.SetPool(ctx, pool)

	// Used to rotate validators so each random operation is applied to a different validator
	validatorCounter := 0

	// Loop through 20 random operations, and check the inflation after each operation
	for i := 0; i < numValidators; i++ {
		pool := keeper.GetPool(ctx)

		// Get inflation before RandomOperation, for comparison later
		previousInflation := pool.Inflation

		// Perform the random operation, and record how validators are modified
		poolMod, validatorMod, _, msg := types.RandomOperation(r)(r, pool, validators[validatorCounter])
		validatorsMod := make([]types.Validator, len(validators))
		copy(validatorsMod[:], validators[:])
		require.Equal(t, numValidators, len(validators), "i %v", validatorCounter)
		require.Equal(t, numValidators, len(validatorsMod), "i %v", validatorCounter)
		validatorsMod[validatorCounter] = validatorMod

		types.AssertInvariants(t, msg,
			pool, validators,
			poolMod, validatorsMod)

		// set pool and validators after the random operation
		pool = poolMod
		keeper.SetPool(ctx, pool)
		validators = validatorsMod

		// Must set inflation here manually, as opposed to most other tests in this suite, where we call keeper.processProvisions(), which updates pool.Inflation
		updatedInflation := keeper.NextInflation(ctx)
		pool.Inflation = updatedInflation
		keeper.SetPool(ctx, pool)

		// Ensure inflation changes as expected when random operations are applied.
		checkInflation(t, pool, previousInflation, updatedInflation, msg)
		validatorCounter++
	}
}

//_________________________________________________________________________________________
////////////////////////////////HELPER FUNCTIONS BELOW/////////////////////////////////////

// Final check on the global pool values for what the total tokens accumulated from each hour of provisions
func checkFinalPoolValues(t *testing.T, pool types.Pool, initialTotalTokens, cumulativeExpProvs sdk.Rat) {
	calculatedTotalTokens := initialTotalTokens.Add(cumulativeExpProvs)
	require.Equal(sdk.RatEq(t, calculatedTotalTokens, pool.TokenSupply()))
}

// Processes provisions are added to the pool correctly every hour
// Returns expected Provisions, expected Inflation, and pool, to help with cumulative calculations back in main Tests
func updateProvisions(t *testing.T, keeper Keeper, pool types.Pool, ctx sdk.Context, hr int) (sdk.Rat, sdk.Rat, types.Pool) {
	expInflation := keeper.NextInflation(ctx)
	expProvisions := expInflation.Mul(pool.TokenSupply()).Quo(hrsPerYrRat)
	startTotalSupply := pool.TokenSupply()
	pool = keeper.ProcessProvisions(ctx)
	keeper.SetPool(ctx, pool)

	//check provisions were added to pool
	require.Equal(sdk.RatEq(t, startTotalSupply.Add(expProvisions), pool.TokenSupply()))

	return expInflation, expProvisions, pool
}

// Deterministic setup of validators and pool
// Allows you to decide how many validators to setup
// Allows you to pick which validators are bonded by adjusting the MaxValidators of params
func setupTestValidators(pool types.Pool, keeper Keeper, ctx sdk.Context, validatorTokens []int64,
	maxValidators uint16) ([]types.Validator, Keeper, types.Pool) {

	params := types.DefaultParams()
	params.MaxValidators = maxValidators
	keeper.SetParams(ctx, params)
	numValidators := len(validatorTokens)
	validators := make([]types.Validator, numValidators)

	for i := 0; i < numValidators; i++ {
		validators[i] = types.NewValidator(Addrs[i], PKs[i], types.Description{})
		validators[i], pool, _ = validators[i].AddTokensFromDel(pool, validatorTokens[i])
		keeper.SetPool(ctx, pool)
		validators[i] = keeper.UpdateValidator(ctx, validators[i]) //will kick out lower power validators. Keep this in mind when setting up the test validators order
		pool = keeper.GetPool(ctx)
	}

	return validators, keeper, pool
}

// Checks that the deterministic validator setup you wanted matches the values in the pool
func checkValidatorSetup(t *testing.T, pool types.Pool, initialTotalTokens, initialBondedTokens int64) {
	require.Equal(t, initialTotalTokens, pool.TokenSupply().RoundInt64(), "%v", pool)
	require.Equal(t, initialBondedTokens, pool.BondedTokens.RoundInt64(), "%v", pool)

	// test initial bonded ratio
	require.True(t, pool.BondedRatio().Equal(sdk.NewRat(initialBondedTokens, initialTotalTokens)), "%v", pool.BondedRatio())
}

// Checks that The inflation will correctly increase or decrease after an update to the pool
// nolint: gocyclo
func checkInflation(t *testing.T, pool types.Pool, previousInflation, updatedInflation sdk.Rat, msg string) {
	inflationChange := updatedInflation.Sub(previousInflation)

	switch {
	//BELOW 67% - Rate of change positive and increasing, while we are between 7% <= and < 20% inflation
	case pool.BondedRatio().LT(sdk.NewRat(67, 100)) && updatedInflation.LT(sdk.NewRat(20, 100)):
		require.Equal(t, true, inflationChange.GT(sdk.ZeroRat()), msg)

	//BELOW 67% - Rate of change should be 0 while inflation continually stays at 20% until we reach 67% bonded ratio
	case pool.BondedRatio().LT(sdk.NewRat(67, 100)) && updatedInflation.Equal(sdk.NewRat(20, 100)):
		if previousInflation.Equal(sdk.NewRat(20, 100)) {
			require.Equal(t, true, inflationChange.IsZero(), msg)

			//This else statement covers the one off case where we first hit 20%, but we still needed a positive ROC to get to 67% bonded ratio (i.e. we went from 19.99999% to 20%)
		} else {
			require.Equal(t, true, inflationChange.GT(sdk.ZeroRat()), msg)
		}

	//ABOVE 67% - Rate of change should be negative while the bond is above 67, and should stay negative until we reach inflation of 7%
	case pool.BondedRatio().GT(sdk.NewRat(67, 100)) && updatedInflation.LT(sdk.NewRat(20, 100)) && updatedInflation.GT(sdk.NewRat(7, 100)):
		require.Equal(t, true, inflationChange.LT(sdk.ZeroRat()), msg)

	//ABOVE 67% - Rate of change should be 0 while inflation continually stays at 7%.
	case pool.BondedRatio().GT(sdk.NewRat(67, 100)) && updatedInflation.Equal(sdk.NewRat(7, 100)):
		if previousInflation.Equal(sdk.NewRat(7, 100)) {
			require.Equal(t, true, inflationChange.IsZero(), msg)

			//This else statement covers the one off case where we first hit 7%, but we still needed a negative ROC to continue to get down to 67%. (i.e. we went from 7.00001% to 7%)
		} else {
			require.Equal(t, true, inflationChange.LT(sdk.ZeroRat()), msg)
		}
	}
}
