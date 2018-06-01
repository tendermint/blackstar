package types

import abci "github.com/tendermint/abci/types"

// initialize application state at genesis
type InitChainer func(ctx Context, req abci.RequestInitChain) abci.ResponseInitChain

// run code before the transactions in a block
type BeginBlocker func(ctx Context, req abci.RequestBeginBlock, res abci.ResponseBeginBlock) abci.ResponseBeginBlock

// run code after the transactions in a block and return updates to the validator set
type EndBlocker func(ctx Context, req abci.RequestEndBlock, res abci.ResponseEndBlock) abci.ResponseEndBlock

// respond to p2p filtering queries from Tendermint
type PeerFilter func(info string) abci.ResponseQuery
