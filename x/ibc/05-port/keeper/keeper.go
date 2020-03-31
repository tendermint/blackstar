package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	"github.com/cosmos/cosmos-sdk/x/ibc/05-port/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/24-host"
)

// Keeper defines the IBC connection keeper
type Keeper struct {
	scopedKeeper capability.ScopedKeeper
	router       map[string]types.IBCModule
}

// NewKeeper creates a new IBC connection Keeper instance
func NewKeeper(sck capability.ScopedKeeper) Keeper {
	return Keeper{
		scopedKeeper: sck,
		router:       make(map[string]types.IBCModule),
	}
}

// isBounded checks a given port ID is already bounded.
func (k Keeper) isBounded(ctx sdk.Context, portID string) bool {
	_, ok := k.scopedKeeper.GetCapability(ctx, types.PortPath(portID))
	return ok
}

// BindPort binds to a port and returns the associated capability.
// Ports must be bound statically when the chain starts in `app.go`.
// The capability must then be passed to a module which will need to pass
// it as an extra parameter when calling functions on the IBC module.
func (k *Keeper) BindPort(ctx sdk.Context, portID string, cbs types.IBCModule) capability.Capability {
	if err := host.DefaultPortIdentifierValidator(portID); err != nil {
		panic(err.Error())
	}

	if k.isBounded(ctx, portID) {
		panic(fmt.Sprintf("port %s is already bound", portID))
	}

	key, err := k.scopedKeeper.NewCapability(ctx, types.PortPath(portID))
	if err != nil {
		panic(err.Error())
	}

	k.router[portID] = cbs

	return key
}

// Authenticate authenticates a capability key against a port ID
// by checking if the memory address of the capability was previously
// generated and bound to the port (provided as a parameter) which the capability
// is being authenticated against.
func (k Keeper) Authenticate(ctx sdk.Context, key capability.Capability, portID string) bool {
	if err := host.DefaultPortIdentifierValidator(portID); err != nil {
		panic(err.Error())
	}

	return k.scopedKeeper.AuthenticateCapability(ctx, key, types.PortPath(portID))
}

// LookupModuleByPort will return the IBCModule along with the capability associated with a given portID
func (k Keeper) LookupModuleByPort(ctx sdk.Context, portID string) (string, capability.Capability, bool) {
	capOwners, ok := k.scopedKeeper.GetOwners(ctx, types.PortPath(portID))
	if !ok {
		return "", nil, false
	}
	if len(capOwners.Owners) < 2 {
		return "", nil, false
	}
	// the module that owns the port, for now, is the first module that isn't "ibc"
	// that owns the module
	module := capOwners.Owners[1].Module
	cap, _ := k.scopedKeeper.GetCapability(ctx, types.PortPath(portID))
	return module, cap, true
}
