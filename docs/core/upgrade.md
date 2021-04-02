<!--
order: 15
-->

# In-Place Store Migrations

Upgrade your app modules smoothly with custom in-place migration logic. {synopsis}

The Cosmos SDK currently has two ways to perform upgrades. The first way is by exporting the entire application state to a JSON file using the `export` CLI command, making changes, and then starting a new binary with the changed JSON file as the genesis file. The second way is by performing upgrades in place, significantly decreasing the time needed to perform upgrades for chains with a larger state. The following guide will provide you with the necessary information in order to setup your application to take advantage of in-place upgrades.

## Genesis State

Each app module's consensus version must be saved to state on the application's genesis. This can be done by adding the following line to the `InitChainer` method in `app.go`

```diff
func (app *MyApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
  ...
+ app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
  ...
}
```

Using this information, the SDK will be able to detect when modules with newer versions are introduced to the app. 

### Consensus Version
The consensus version is defined on each app module by the module developer. It serves as the breaking change version of the module. This number is how the SDK identifies which modules to upgrade. For example, if the bank module was version 2, and an upgrade introduces bank module 3, the SDK will know to upgrade the bank module.

### Version Map
The version map is a mapping of module names to consensus versions. The map is persisted to state for use during in-place migrations. When migrations finish, the updated version map is persisted to state. 

## Upgrade Handlers

Upgrades utilize an `UpgradeHandler` to facilitate migrations. `UpgradeHandler`s are functions implemented by the app developer that conform to the following function signature. These functions utiltize a `VersionMap` containing all the module versions to determine which modules need upgrading.

```golang
type UpgradeHandler func(ctx sdk.Context, plan Plan, fromVM VersionMap) (VersionMap, error)
```

Inside these functions, you should perform any upgrade logic you wish to include in the provided `plan`. All upgrade handler functions should end with the following line of code:

```golang
  return app.mm.RunMigrations(ctx, cfg, fromVM)
```

## Running Migrations

Migrations are run inside of an `UpgradeHandler` via `app.RunMigrations(ctx, cfg, vm)`. As described above, `UpgradeHandler`s are functions which describe the functionality to occur during an upgrade. The `RunMigration` function will loop through the `VersionMap` argument, and run the migration scripts for any versions that are less than the new binary's app module versions. Once the migrations are finished, a new `VersionMap` will be returned to persist the upgraded module versions to state.

```golang
cfg := module.NewConfigurator(...)
app.UpgradeKeeper.SetUpgradeHandler("my-plan", func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {

    // ...
    // do upgrade logic
    // ...

    // RunMigrations returns the VersionMap
    // with the updated module ConsensusVersions
    return app.RunMigrations(ctx, vm)
})
```

To learn more about configuring migration scripts, refer to this (guide)[../building-modules/upgrade.md].

## Adding New Modules In Upgrades

Entirely new modules can be introduced to the application during an upgrade. The SDK recognizes new modules during upgrades and will call the corresponding module's `InitGenesis` function to setup its initial state. This can be skipped if the module does not require any inital state. Otherwise, it is important to implement `InitGenesis` for new modules to successfully upgrade your application without error.

In the scenario where your application does not need any inital state via `InitGenesis`, you must take extra steps to ensure `InitGenesis` is skipped to avoid errors. To do so, you simply need to update the value of the module version in the `VersionMap` in the `UpgradeHandler`. 

```go
// Foo is a new module being introduced
// in this upgrade plan
import foo "github.com/my/module/foo"

app.UpgradeKeeper.SetUpgradeHandler("my-plan", func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap)  (module.VersionMap, error) {
    // We make sure to set foo's version to the latest ConsensusVersion in the VersionMap.
    // This will skip running InitGenesis on Foo
    vm["foo"] = foo.AppModule{}.ConsensusVersion()

    return app.mm.RunMigrations(ctx, cfg, vm)
})
```

Using a similar method, you can also run InitGenesis on your new module with a custom genesis state:

```go
import foo "github.com/my/module/foo"

app.UpgradeKeeper.SetUpgradeHandler("my-plan", func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap)  (module.VersionMap, error) {
    vm["foo"] = foo.AppModule{}.ConsensusVersion()

    // Run custom InitGenesis for foo
    app.mm["foo"].InitGenesis(ctx, app.appCodec, myCustomGenesisState)

    return app.mm.RunMigrations(ctx, cfg, vm)
})
```