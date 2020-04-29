package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/evidence/exported"
)

// Codec defines the interface required to serialize evidence
type Codec interface {
	codec.Marshaler

	MarshalEvidence(exported.Evidence) ([]byte, error)
	UnmarshalEvidence([]byte) (exported.Evidence, error)
	MarshalEvidenceJSON(exported.Evidence) ([]byte, error)
	UnmarshalEvidenceJSON([]byte) (exported.Evidence, error)
}

// RegisterCodec registers all the necessary types and interfaces for the
// evidence module.
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*exported.Evidence)(nil), nil)
	cdc.RegisterConcrete(MsgSubmitEvidence{}, "cosmos-sdk/MsgSubmitEvidence", nil)
	cdc.RegisterConcrete(&Equivocation{}, "cosmos-sdk/Equivocation", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgSubmitEvidence{})
	registry.RegisterInterface(
		"cosmos_sdk.evidence.v1.Evidence",
		(*exported.Evidence)(nil),
		&Equivocation{},
	)
}

var (
	amino = codec.New()

	// ModuleCdc references the global x/evidence module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/evidence and
	// defined at the application level.
	ModuleCdc = codec.NewHybridCodec(amino, types.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(amino)
	codec.RegisterCrypto(amino)
	amino.Seal()
}
