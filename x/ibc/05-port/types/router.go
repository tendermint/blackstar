package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Router struct {
	routes map[string]IBCModule
	sealed bool
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]IBCModule),
	}
}

// Seal prevents the Router from any subsequent route handlers to be registered.
// Seal will panic if called more than once.
func (rtr *Router) Seal() {
	if rtr.sealed {
		panic("Router already sealed")
	}
	rtr.sealed = true
}

// Sealed returns a boolean signifying if the Router is sealed or not.
func (rtr Router) Sealed() bool {
	return rtr.sealed
}

// AddRoute adds IBCModule for a given module name. It returns the Router
// so AddRoute calls can be linked. It will panic if the Router is sealed.
func (rtr *Router) AddRoute(module string, cbs IBCModule) *Router {
	if rtr.sealed {
		panic(fmt.Sprintf("Router sealed; cannot register %s route callbacks", module))
	}
	if !sdk.IsAlphaNumeric(module) {
		panic("route expressions can only contain alphanumeric characters")
	}
	if rtr.HasRoute(module) {
		panic(fmt.Sprintf("route %s has already been registered", module))
	}

	rtr.routes[module] = cbs
	return rtr
}

// HasRoute returns true if the Router has a module registered or false otherwise.
func (rtr *Router) HasRoute(module string) bool {
	return rtr.routes[module] != nil
}

// GetRoute returns a IBCModule for a given module.
func (rtr *Router) GetRoute(module string) IBCModule {
	if !rtr.HasRoute(module) {
		panic(fmt.Sprintf("route does not exist for module %s", module))
	}
	return rtr.routes[module]
}
