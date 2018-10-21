package simplestake

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Amino) {
	cdc.RegisterConcrete(MsgBond{}, "simplestake/BondMsg", nil)
	cdc.RegisterConcrete(MsgUnbond{}, "simplestake/UnbondMsg", nil)
}
