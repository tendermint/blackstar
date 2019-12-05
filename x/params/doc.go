/*
Package params provides a namespaced module parameter store.

There are two core components, Keeper and Subspace. Subspace is an isolated
namespace for a parameter store, where keys are prefixed by pre-configured
subspace names which modules provide. The Keeper has a permission to access all
existing subspaces.

Subspace can be used by the individual keepers, who needs a private parameter store
that the other keeper cannot modify. Keeper can be used by the Governance keeper,
who need to modify any parameter in case of the proposal passes.

Basic Usage:

1. Declare constant module parameter keys and the globally unique Subspace name:

	const (
		ModuleSubspace = "mymodule"
	)

	const (
		KeyParameter1 = "myparameter1"
		KeyParameter2 = "myparameter2"
	)

2. Create a concrete parameter struct and define the validation functions:

	type MyParams struct {
		MyParam1 int64 `json:"my_param1" yaml:"my_param1"`
		MyParam2 bool `json:"my_param2" yaml:"my_param2"`
	}

	func validateMyParam1(i interface{}) error {
		_, ok := i.(int64)
		if !ok {
			return fmt.Errorf("invalid parameter type: %T", i)
		}

		// validate (if necessary)...

		return nil
	}

	func validateMyParam2(i interface{}) error {
		_, ok := i.(bool)
		if !ok {
			return fmt.Errorf("invalid parameter type: %T", i)
		}

		// validate (if necessary)...

		return nil
	}

3. Implement the params.ParamSet interface:

	func (p *MyParams) ParamSetPairs() params.ParamSetPairs {
		return params.ParamSetPairs{
			{KeyParameter1, &p.MyParam1, validateMyParam1},
			{KeyParameter2, &p.MyParam2, validateMyParam2},
		}
	}

	func paramKeyTable() params.KeyTable {
		return params.NewKeyTable().RegisterParamSet(&MyParams{})
	}

4. Have the module accept a Subspace in the constructor and set the KeyTable (if necessary):

	func NewKeeper(..., paramSpace params.Subspace, ...) Keeper {
		// set KeyTable if it has not already been set
		if !paramSpace.HasKeyTable() {
			paramSpace = paramSpace.WithKeyTable(paramKeyTable())
		}

		return Keeper {
			// ...
			paramSpace: paramSpace,
		}
	}

Now we have access to the module's paramters that are namespaced using the keys defined:

	func InitGenesis(ctx sdk.Context, k Keeper, gs GenesisState) {
		// ...
		k.SetParams(ctx, gs.Params)
	}

	func (k Keeper) SetParams(ctx sdk.Context, params Params) {
		k.paramSpace.SetParamSet(ctx, &params)
	}

	func (k Keeper) GetParams(ctx sdk.Context) (params Params) {
		k.paramSpace.GetParamSet(ctx, &params)
		return params
	}

	func (k Keeper) MyParam1(ctx sdk.Context) (res int64) {
		k.paramSpace.Get(ctx, KeyParameter1, &res)
		return res
	}

	func (k Keeper) MyParam2(ctx sdk.Context) (res bool) {
		k.paramSpace.Get(ctx, KeyParameter2, &res)
		return res
	}

NOTE: Any call to SetParamSet will panic or any call to Update will error if any
given parameter value is invalid based on the registered value validation function.
*/
package params
