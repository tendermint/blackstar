package keeper

import (
	"container/list"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/stake/types"
)

// Cache the amino decoding of validators, as it can be the case that repeated slashing calls
// cause many calls to GetValidator, which were shown to throttle the state machine in our
// simulation. Note this is quite biased though, as the simulator does more slashes than a
// live chain should, however we require the slashing to be fast as noone pays gas for it.
type cachedValidator struct {
	val        types.Validator
	marshalled string // marshalled amino bytes for the validator object (not operator address)
}

// validatorCache-key: validator amino bytes
var validatorCache = make(map[string]cachedValidator, 500)
var validatorCacheList = list.New()

// get a single validator
func (k Keeper) GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(GetValidatorKey(addr))
	if value == nil {
		return validator, false
	}

	// If these amino encoded bytes are in the cache, return the cached validator
	strValue := string(value)
	if val, ok := validatorCache[strValue]; ok {
		valToReturn := val.val
		// Doesn't mutate the cache's value
		valToReturn.OperatorAddr = addr
		return valToReturn, true
	}

	// amino bytes weren't found in cache, so amino unmarshal and add it to the cache
	validator = types.MustUnmarshalValidator(k.cdc, addr, value)
	cachedVal := cachedValidator{validator, strValue}
	validatorCache[strValue] = cachedValidator{validator, strValue}
	validatorCacheList.PushBack(cachedVal)

	// if the cache is too big, pop off the last element from it
	if validatorCacheList.Len() > 500 {
		valToRemove := validatorCacheList.Remove(validatorCacheList.Front()).(cachedValidator)
		delete(validatorCache, valToRemove.marshalled)
	}

	validator = types.MustUnmarshalValidator(k.cdc, addr, value)
	return validator, true
}

func (k Keeper) mustGetValidator(ctx sdk.Context, addr sdk.ValAddress) types.Validator {
	validator, found := k.GetValidator(ctx, addr)
	if !found {
		panic(fmt.Sprintf("validator record not found for address: %X\n", addr))
	}
	return validator
}

// get a single validator by consensus address
func (k Keeper) GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)
	opAddr := store.Get(GetValidatorByConsAddrKey(consAddr))
	if opAddr == nil {
		return validator, false
	}
	return k.GetValidator(ctx, opAddr)
}

func (k Keeper) mustGetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) types.Validator {
	validator, found := k.GetValidatorByConsAddr(ctx, consAddr)
	if !found {
		panic(fmt.Errorf("validator with consensus-Address %s not found", consAddr))
	}
	return validator
}

//___________________________________________________________________________

// set the main record holding validator details
func (k Keeper) SetValidator(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalValidator(k.cdc, validator)
	store.Set(GetValidatorKey(validator.OperatorAddr), bz)
}

// validator index
func (k Keeper) SetValidatorByConsAddr(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	consAddr := sdk.ConsAddress(validator.ConsPubKey.Address())
	store.Set(GetValidatorByConsAddrKey(consAddr), validator.OperatorAddr)
}

// validator index
func (k Keeper) SetValidatorByPowerIndex(ctx sdk.Context, validator types.Validator, pool types.Pool) {
	// jailed validators are not kept in the power index
	if validator.Jailed {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(GetBondedValidatorsByPowerIndexKey(validator, pool), validator.OperatorAddr)
}

// validator index
func (k Keeper) DeleteValidatorByPowerIndex(ctx sdk.Context, validator types.Validator, pool types.Pool) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(GetBondedValidatorsByPowerIndexKey(validator, pool))
}

// validator index
func (k Keeper) SetNewValidatorByPowerIndex(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	pool := k.GetPool(ctx)
	store.Set(GetBondedValidatorsByPowerIndexKey(validator, pool), validator.OperatorAddr)
}

//___________________________________________________________________________

// Update the tokens of an existing validator, update the validators power index key
func (k Keeper) AddValidatorTokensAndShares(ctx sdk.Context, validator types.Validator,
	tokensToAdd sdk.Int) (valOut types.Validator, addedShares sdk.Dec) {

	pool := k.GetPool(ctx)
	k.DeleteValidatorByPowerIndex(ctx, validator, pool)
	validator, pool, addedShares = validator.AddTokensFromDel(pool, tokensToAdd)
	// increment the intra-tx counter
	// in case of a conflict, the validator which least recently changed power takes precedence
	counter := k.GetIntraTxCounter(ctx)
	validator.BondIntraTxCounter = counter
	k.SetIntraTxCounter(ctx, counter+1)
	k.SetValidator(ctx, validator)
	k.SetPool(ctx, pool)
	k.SetValidatorByPowerIndex(ctx, validator, pool)
	return validator, addedShares
}

// Update the tokens of an existing validator, update the validators power index key
func (k Keeper) RemoveValidatorTokensAndShares(ctx sdk.Context, validator types.Validator,
	sharesToRemove sdk.Dec) (valOut types.Validator, removedTokens sdk.Dec) {

	pool := k.GetPool(ctx)
	k.DeleteValidatorByPowerIndex(ctx, validator, pool)
	validator, pool, removedTokens = validator.RemoveDelShares(pool, sharesToRemove)
	k.SetValidator(ctx, validator)
	k.SetPool(ctx, pool)
	k.SetValidatorByPowerIndex(ctx, validator, pool)
	return validator, removedTokens
}

// Update the tokens of an existing validator, update the validators power index key
func (k Keeper) RemoveValidatorTokens(ctx sdk.Context, validator types.Validator, tokensToRemove sdk.Dec) types.Validator {
	pool := k.GetPool(ctx)
	k.DeleteValidatorByPowerIndex(ctx, validator, pool)
	validator, pool = validator.RemoveTokens(pool, tokensToRemove)
	k.SetValidator(ctx, validator)
	k.SetPool(ctx, pool)
	k.SetValidatorByPowerIndex(ctx, validator, pool)
	return validator
}

// UpdateValidatorCommission attempts to update a validator's commission rate.
// An error is returned if the new commission rate is invalid.
func (k Keeper) UpdateValidatorCommission(ctx sdk.Context, validator types.Validator, newRate sdk.Dec) (types.Commission, sdk.Error) {
	commission := validator.Commission
	blockTime := ctx.BlockHeader().Time

	if err := commission.ValidateNewRate(newRate, blockTime); err != nil {
		return commission, err
	}

	commission.Rate = newRate
	commission.UpdateTime = blockTime

	return commission, nil
}

// remove the validator record and associated indexes
func (k Keeper) RemoveValidator(ctx sdk.Context, address sdk.ValAddress) {

	// first retrieve the old validator record
	validator, found := k.GetValidator(ctx, address)
	if !found {
		return
	}

	// delete the old validator record
	store := ctx.KVStore(k.storeKey)
	pool := k.GetPool(ctx)
	store.Delete(GetValidatorKey(address))
	store.Delete(GetValidatorByConsAddrKey(sdk.ConsAddress(validator.ConsPubKey.Address())))
	store.Delete(GetBondedValidatorsByPowerIndexKey(validator, pool))

}

//___________________________________________________________________________
// get groups of validators

// get the set of all validators with no limits, used during genesis dump
func (k Keeper) GetAllValidators(ctx sdk.Context) (validators []types.Validator) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, ValidatorsKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		addr := iterator.Key()[1:]
		validator := types.MustUnmarshalValidator(k.cdc, addr, iterator.Value())
		validators = append(validators, validator)
	}
	return validators
}

// return a given amount of all the validators
func (k Keeper) GetValidators(ctx sdk.Context, maxRetrieve uint16) (validators []types.Validator) {
	store := ctx.KVStore(k.storeKey)
	validators = make([]types.Validator, maxRetrieve)

	iterator := sdk.KVStorePrefixIterator(store, ValidatorsKey)
	defer iterator.Close()

	i := 0
	for ; iterator.Valid() && i < int(maxRetrieve); iterator.Next() {
		addr := iterator.Key()[1:]
		validator := types.MustUnmarshalValidator(k.cdc, addr, iterator.Value())
		validators[i] = validator
		i++
	}
	return validators[:i] // trim if the array length < maxRetrieve
}

// get the group of the bonded validators
func (k Keeper) GetValidatorsBonded(ctx sdk.Context) (validators []types.Validator) {
	store := ctx.KVStore(k.storeKey)

	// add the actual validator power sorted store
	maxValidators := k.GetParams(ctx).MaxValidators
	validators = make([]types.Validator, maxValidators)

	iterator := sdk.KVStorePrefixIterator(store, ValidatorsBondedIndexKey)
	defer iterator.Close()

	i := 0
	for ; iterator.Valid(); iterator.Next() {

		// sanity check
		if i > int(maxValidators-1) {
			panic("maxValidators is less than the number of records in ValidatorsBonded Store, store should have been updated")
		}
		address := GetAddressFromValBondedIndexKey(iterator.Key())
		validator := k.mustGetValidator(ctx, address)

		validators[i] = validator
		i++
	}
	return validators[:i] // trim
}

// get the group of bonded validators sorted by power-rank
func (k Keeper) GetBondedValidatorsByPower(ctx sdk.Context) []types.Validator {
	store := ctx.KVStore(k.storeKey)
	maxValidators := k.GetParams(ctx).MaxValidators
	validators := make([]types.Validator, maxValidators)

	iterator := sdk.KVStoreReversePrefixIterator(store, ValidatorsByPowerIndexKey)
	defer iterator.Close()

	i := 0
	for ; iterator.Valid() && i < int(maxValidators); iterator.Next() {
		address := iterator.Value()
		validator := k.mustGetValidator(ctx, address)

		if validator.Status == sdk.Bonded {
			validators[i] = validator
			i++
		}
	}
	return validators[:i] // trim
}
