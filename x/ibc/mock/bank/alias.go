package mockbank

import (
	"github.com/cosmos/cosmos-sdk/x/ibc/mock/bank/internal/keeper"
	"github.com/cosmos/cosmos-sdk/x/ibc/mock/bank/internal/types"
)

// nolint
type (
	MsgTransfer = types.MsgTransfer
	Keeper      = keeper.Keeper
)

// nolint
var (
	ModuleName   = types.ModuleName
	StoreKey     = types.StoreKey
	TStoreKey    = types.TStoreKey
	QuerierRoute = types.QuerierRoute
	RouterKey    = types.RouterKey

	RegisterCdc = types.RegisterCodec
)
