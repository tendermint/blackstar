package tx

import (
	"fmt"

	"github.com/golang/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
)

// DefaultTxEncoder returns a default protobuf TxEncoder using the provided Marshaler
func DefaultTxEncoder() types.TxEncoder {
	return func(tx types.Tx) ([]byte, error) {
		wrapper, ok := tx.(*builder)
		if !ok {
			return nil, fmt.Errorf("expected %T, got %T", &builder{}, tx)
		}

		raw := &txtypes.TxRaw{
			BodyBytes:     wrapper.GetBodyBytes(),
			AuthInfoBytes: wrapper.GetAuthInfoBytes(),
			Signatures:    wrapper.tx.Signatures,
		}

		return proto.Marshal(raw)
	}
}

// DefaultTxEncoder returns a default protobuf JSON TxEncoder using the provided Marshaler
func DefaultJSONTxEncoder() types.TxEncoder {
	return func(tx types.Tx) ([]byte, error) {
		wrapper, ok := tx.(*builder)
		if !ok {
			return nil, fmt.Errorf("expected %T, got %T", &builder{}, tx)
		}

		return codec.ProtoMarshalJSON(wrapper.tx)
	}
}
