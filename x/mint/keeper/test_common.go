package keeper

import (
	"os"
	"testing"
	"time"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/stretchr/testify/require"

	dbm "github.com/tendermint/tendermint/libs/db"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"

	"github.com/cosmos/cosmos-sdk/x/mint/types"
)

type testInput struct {
	ctx        sdk.Context
	cdc        *codec.Codec
	mintKeeper Keeper
}

func createTestCodec() *codec.Codec {
	cdc := codec.New()
	codec.RegisterCrypto(cdc)
	return cdc
}

func newTestInput(t *testing.T) testInput {
	cdc := createTestCodec()
	db := dbm.NewMemDB()

	keyAcc := sdk.NewKVStoreKey(auth.StoreKey)
	keyStaking := sdk.NewKVStoreKey(staking.StoreKey)
	tkeyStaking := sdk.NewTransientStoreKey(staking.TStoreKey)
	keyParams := sdk.NewKVStoreKey(params.StoreKey)
	tkeyParams := sdk.NewTransientStoreKey(params.TStoreKey)
	keyFeeCollection := sdk.NewKVStoreKey(auth.FeeStoreKey)
	keySupply := sdk.NewKVStoreKey(supply.StoreKey)
	keyMint := sdk.NewKVStoreKey(StoreKey)

	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyAcc, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tkeyStaking, sdk.StoreTypeTransient, nil)
	ms.MountStoreWithDB(keyStaking, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyFeeCollection, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyMint, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, db)
	err := ms.LoadLatestVersion()
	require.Nil(t, err)

	paramsKeeper := params.NewKeeper(cdc, keyParams, tkeyParams)
	feeCollectionKeeper := auth.NewFeeCollectionKeeper(cdc, keyFeeCollection)
	accountKeeper := auth.NewAccountKeeper(cdc, keyAcc, paramsKeeper.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	bankKeeper := bank.NewBaseKeeper(accountKeeper, paramsKeeper.Subspace(bank.DefaultParamspace), bank.DefaultCodespace)
	stakingKeeper := staking.NewKeeper(
		cdc, keyStaking, tkeyStaking, bankKeeper, paramsKeeper.Subspace(staking.DefaultParamspace), staking.DefaultCodespace,
	)
	supplyKeeper := supply.NewKeeper(cdc, keySupply)

	mintKeeper := NewKeeper(
		cdc, keyMint, paramsKeeper.Subspace(DefaultParamspace), supplyKeeper, &stakingKeeper, feeCollectionKeeper,
	)

	ctx := sdk.NewContext(ms, abci.Header{Time: time.Unix(0, 0)}, false, log.NewTMLogger(os.Stdout))

	mintKeeper.SetParams(ctx, types.DefaultParams())
	mintKeeper.SetMinter(ctx, types.DefaultInitialMinter())

	return testInput{ctx, cdc, mintKeeper}
}
