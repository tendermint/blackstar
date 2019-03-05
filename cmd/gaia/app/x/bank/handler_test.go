package bank

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
)

var (
	addrs = []sdk.AccAddress{
		sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()),
		sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()),
		sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()),
	}

	testDenom = "test"
	initAmt   = sdk.NewInt(10000)
)

func newTestCodec() *codec.Codec {
	cdc := codec.New()

	bank.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	return cdc
}

func createTestInput(t *testing.T) (sdk.Context, auth.AccountKeeper, bank.Keeper) {
	keyAcc := sdk.NewKVStoreKey(auth.StoreKey)
	keyParams := sdk.NewKVStoreKey(params.StoreKey)
	tKeyParams := sdk.NewTransientStoreKey(params.TStoreKey)

	cdc := newTestCodec()
	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ctx := sdk.NewContext(ms, abci.Header{Time: time.Now().UTC()}, false, log.NewNopLogger())

	ms.MountStoreWithDB(keyAcc, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tKeyParams, sdk.StoreTypeTransient, db)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)

	require.NoError(t, ms.LoadLatestVersion())

	paramsKeeper := params.NewKeeper(cdc, keyParams, tKeyParams)
	authKeeper := auth.NewAccountKeeper(
		cdc,
		keyAcc,
		paramsKeeper.Subspace(auth.DefaultParamspace),
		auth.ProtoBaseAccount,
	)

	bankKeeper := bank.NewBaseKeeper(
		authKeeper,
		paramsKeeper.Subspace(bank.DefaultParamspace),
		bank.DefaultCodespace,
	)

	for _, addr := range addrs {
		_, _, err := bankKeeper.AddCoins(ctx, addr, sdk.Coins{sdk.NewCoin(testDenom, initAmt)})
		require.NoError(t, err)
	}

	return ctx, authKeeper, bankKeeper
}

func TestHandlerMsgSendTransfersDisabled(t *testing.T) {
	ctx, ak, bk := createTestInput(t)
	bk.SetSendEnabled(ctx, false)

	handler := NewHandler(bk)
	amt := sdk.NewInt(1000)
	msg := bank.NewMsgSend(addrs[0], addrs[1], sdk.Coins{sdk.NewCoin(testDenom, amt)})

	res := handler(ctx, msg)
	require.False(t, res.IsOK())

	from := ak.GetAccount(ctx, addrs[0])
	require.Equal(t, from.GetCoins(), sdk.Coins{sdk.NewCoin(testDenom, initAmt)})

	to := ak.GetAccount(ctx, addrs[1])
	require.Equal(t, to.GetCoins(), sdk.Coins{sdk.NewCoin(testDenom, initAmt)})
}

func TestHandlerMsgSendTransfersEnabled(t *testing.T) {
	ctx, ak, bk := createTestInput(t)
	bk.SetSendEnabled(ctx, true)

	handler := NewHandler(bk)
	amt := sdk.NewInt(1000)
	msg := bank.NewMsgSend(addrs[0], addrs[1], sdk.Coins{sdk.NewCoin(testDenom, amt)})

	res := handler(ctx, msg)
	require.True(t, res.IsOK())

	from := ak.GetAccount(ctx, addrs[0])
	balance := initAmt.Sub(amt)
	require.Equal(t, from.GetCoins(), sdk.Coins{sdk.NewCoin(testDenom, balance)})

	to := ak.GetAccount(ctx, addrs[1])
	balance = initAmt.Add(amt)
	require.Equal(t, to.GetCoins(), sdk.Coins{sdk.NewCoin(testDenom, balance)})
}
