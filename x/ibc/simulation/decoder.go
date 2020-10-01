package simulation

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/kv"
	connectionsim "github.com/cosmos/cosmos-sdk/x/ibc/03-connection/simulation"
	channelsim "github.com/cosmos/cosmos-sdk/x/ibc/04-channel/simulation"
	host "github.com/cosmos/cosmos-sdk/x/ibc/24-host"
	clientsim "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/simulation"
	"github.com/cosmos/cosmos-sdk/x/ibc/keeper"
)

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding ibc type.
func NewDecodeStore(k keeper.Keeper) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		if res, found := clientsim.NewDecodeStore(k.ClientKeeper, kvA, kvB); found {
			return res
		}

		if res, found := connectionsim.NewDecodeStore(k.Codec(), kvA, kvB); found {
			return res
		}

		if res, found := channelsim.NewDecodeStore(k.Codec(), kvA, kvB); found {
			return res
		}

		panic(fmt.Sprintf("invalid %s key prefix: %s", host.ModuleName, string(kvA.Key)))
	}
}
