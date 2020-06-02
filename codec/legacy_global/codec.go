package legacy_global

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codec2 "github.com/cosmos/cosmos-sdk/crypto/codec"
)

// Deprecated: Cdc defines a global generic sealed Amino codec to be used throughout sdk. It
// has all Tendermint crypto and evidence types registered.
//
// TODO: remove this global.
var Cdc *codec.Codec

func init() {
	Cdc = codec.New()
	codec2.RegisterCrypto(Cdc)
	codec.RegisterEvidences(Cdc)
	Cdc.Seal()
}
