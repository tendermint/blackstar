/*
# Simulation App

The `SimApp` defines an application used for running extensive simulation
testing suites. It contains all core modules, including governance, staking,
slashing, and distribution.

Simulation is executed with various inputs including the number of blocks to
simulate, the block size, whether the app should commit or not, the invariant
checking period, and a seed which is used as a source of pseudo-randomness.

In addition to the various inputs, simulation runs mainly in three modes:

* Completely random where the initial state, module parameters and simulation
parameters are pseudo-randomly generated.
* From a genesis file where the initial state and the module parameters are defined.
This mode is helpful for running simulations on a known state such as a live
network export where a new (mostly likely breaking) version of the application
needs to be tested.
* From a params file where the initial state is pseudo-randomly generated but the
module and simulation parameters can be provided manually. This allows for a more
controlled and deterministic simulation setup while allowing the state space to
still be pseudo-randomly simulated.

The simulation test suite also supports testing determinism and import/export
functionality.

## Randomness

Currently, simulation uses a single seed (integer) as a source for a PRNG by
which all random operations are executed from. Any call to the PRNG changes all
future operations as the internal state of the PRNG is modified. For example,
if a new message type is created and needs to be simulated, the new introduced
PRNG call will change all subsequent operations.

This may can often be problematic when testing fixes to simulation faults. One
current solution to this is to use a params file as mentioned above. In the future
the simulation suite is expected to support a series of PRNGs that can be used
uniquely per module and simulation component so that they will not effect each
others state execution outcome.

## Usage

To execute a completely pseudo-random simulation:

```shell
$ go test -mod=readonly github.com/cosmos/cosmos-sdk/simapp \
  -run=TestFullAppSimulation \
  -SimulationEnabled=true \
  -SimulationNumBlocks=100 \
  -SimulationBlockSize=200 \
  -SimulationCommit=true \
  -SimulationSeed=99 \
  -SimulationPeriod=5 \
  -v -timeout 24h
```

To execute simulation from a genesis file:

```shell
$ go test -mod=readonly github.com/cosmos/cosmos-sdk/simapp \
  -run=TestFullAppSimulation \
  -SimulationEnabled=true \
  -SimulationNumBlocks=100 \
  -SimulationBlockSize=200 \
  -SimulationCommit=true \
  -SimulationSeed=99 \
  -SimulationPeriod=5 \
  -SimulationGenesis=/path/to/genesis.json \
  -v -timeout 24h
```

To execute simulation from a params file:

```shell
$ go test -mod=readonly github.com/cosmos/cosmos-sdk/simapp \
  -run=TestFullAppSimulation \
  -SimulationEnabled=true \
  -SimulationNumBlocks=100 \
  -SimulationBlockSize=200 \
  -SimulationCommit=true \
  -SimulationSeed=99 \
  -SimulationPeriod=5 \
  -SimulationParams=/path/to/params.json \
  -v -timeout 24h
```

## Params

Params that are provided to simulation from a JSON file are used to used to set
both module parameters and simulation parameters. See `sim_test.go` for the full
set of parameters that can be provided.
*/
package simapp
