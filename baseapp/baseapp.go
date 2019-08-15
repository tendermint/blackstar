package baseapp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"runtime/debug"
	"strings"

	"github.com/gogo/protobuf/proto"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// Check a transaction
	runTxModeCheck runTxMode = iota
	// Simulate a transaction
	runTxModeSimulate runTxMode = iota
	// Deliver a transaction
	runTxModeDeliver runTxMode = iota

	// MainStoreKey is the string representation of the main store
	MainStoreKey = "main"
)

var (
	_ abci.Application = (*BaseApp)(nil)

	// mainConsensusParamsKey defines a key to store the consensus params in the
	// main store.
	mainConsensusParamsKey = []byte("consensus_params")
)

type (
	// Enum mode for app.runTx
	runTxMode uint8

	// StoreLoader defines a customizable function to control how we load the CommitMultiStore
	// from disk. This is useful for state migration, when loading a datastore written with
	// an older version of the software. In particular, if a module changed the substore key name
	// (or removed a substore) between two versions of the software.
	StoreLoader func(ms sdk.CommitMultiStore) error
)

// BaseApp reflects the ABCI application implementation.
type BaseApp struct {
	// initialized on creation
	logger      log.Logger
	name        string               // application name from abci.Info
	db          dbm.DB               // common DB backend
	cms         sdk.CommitMultiStore // Main (uncached) state
	storeLoader StoreLoader          // function to handle store loading, may be overridden with SetStoreLoader()
	router      sdk.Router           // handle any kind of message
	queryRouter sdk.QueryRouter      // router for redirecting query calls
	txDecoder   sdk.TxDecoder        // unmarshal []byte into sdk.Tx

	// set upon LoadVersion or LoadLatestVersion.
	baseKey *sdk.KVStoreKey // Main KVStore in cms

	anteHandler    sdk.AnteHandler  // ante handler for fee and auth
	initChainer    sdk.InitChainer  // initialize state with validators and state blob
	beginBlocker   sdk.BeginBlocker // logic to run before any txs
	endBlocker     sdk.EndBlocker   // logic to run after all txs, and to determine valset changes
	addrPeerFilter sdk.PeerFilter   // filter peers by address and port
	idPeerFilter   sdk.PeerFilter   // filter peers by node ID
	fauxMerkleMode bool             // if true, IAVL MountStores uses MountStoresDB for simulation speed.

	// --------------------
	// Volatile state
	// checkState is set on initialization and reset on Commit.
	// deliverState is set in InitChain and BeginBlock and cleared on Commit.
	// See methods setCheckState and setDeliverState.
	checkState   *state          // for CheckTx
	deliverState *state          // for DeliverTx
	voteInfos    []abci.VoteInfo // absent validators from begin block

	// consensus params
	// TODO: Move this in the future to baseapp param store on main store.
	consensusParams *abci.ConsensusParams

	// The minimum gas prices a validator is willing to accept for processing a
	// transaction. This is mainly used for DoS and spam prevention.
	minGasPrices sdk.DecCoins

	// flag for sealing options and parameters to a BaseApp
	sealed bool

	// height at which to halt the chain and gracefully shutdown
	haltHeight uint64

	// application's version string
	appVersion string
}

// NewBaseApp returns a reference to an initialized BaseApp. It accepts a
// variadic number of option functions, which act on the BaseApp to set
// configuration choices.
//
// NOTE: The db is used to store the version number for now.
func NewBaseApp(
	name string, logger log.Logger, db dbm.DB, txDecoder sdk.TxDecoder, options ...func(*BaseApp),
) *BaseApp {

	app := &BaseApp{
		logger:         logger,
		name:           name,
		db:             db,
		cms:            store.NewCommitMultiStore(db),
		storeLoader:    DefaultStoreLoader,
		router:         NewRouter(),
		queryRouter:    NewQueryRouter(),
		txDecoder:      txDecoder,
		fauxMerkleMode: false,
	}
	for _, option := range options {
		option(app)
	}

	return app
}

// Name returns the name of the BaseApp.
func (app *BaseApp) Name() string {
	return app.name
}

// AppVersion returns the application's version string.
func (app *BaseApp) AppVersion() string {
	return app.appVersion
}

// Logger returns the logger of the BaseApp.
func (app *BaseApp) Logger() log.Logger {
	return app.logger
}

// MountStores mounts all IAVL or DB stores to the provided keys in the BaseApp
// multistore.
func (app *BaseApp) MountStores(keys ...sdk.StoreKey) {
	for _, key := range keys {
		switch key.(type) {
		case *sdk.KVStoreKey:
			if !app.fauxMerkleMode {
				app.MountStore(key, sdk.StoreTypeIAVL)
			} else {
				// StoreTypeDB doesn't do anything upon commit, and it doesn't
				// retain history, but it's useful for faster simulation.
				app.MountStore(key, sdk.StoreTypeDB)
			}

		case *sdk.TransientStoreKey:
			app.MountStore(key, sdk.StoreTypeTransient)

		default:
			panic("Unrecognized store key type " + reflect.TypeOf(key).Name())
		}
	}
}

// MountStores mounts all IAVL or DB stores to the provided keys in the BaseApp
// multistore.
func (app *BaseApp) MountKVStores(keys map[string]*sdk.KVStoreKey) {
	for _, key := range keys {
		if !app.fauxMerkleMode {
			app.MountStore(key, sdk.StoreTypeIAVL)
		} else {
			// StoreTypeDB doesn't do anything upon commit, and it doesn't
			// retain history, but it's useful for faster simulation.
			app.MountStore(key, sdk.StoreTypeDB)
		}
	}
}

// MountStores mounts all IAVL or DB stores to the provided keys in the BaseApp
// multistore.
func (app *BaseApp) MountTransientStores(keys map[string]*sdk.TransientStoreKey) {
	for _, key := range keys {
		app.MountStore(key, sdk.StoreTypeTransient)
	}
}

// MountStoreWithDB mounts a store to the provided key in the BaseApp
// multistore, using a specified DB.
func (app *BaseApp) MountStoreWithDB(key sdk.StoreKey, typ sdk.StoreType, db dbm.DB) {
	app.cms.MountStoreWithDB(key, typ, db)
}

// MountStore mounts a store to the provided key in the BaseApp multistore,
// using the default DB.
func (app *BaseApp) MountStore(key sdk.StoreKey, typ sdk.StoreType) {
	app.cms.MountStoreWithDB(key, typ, nil)
}

// LoadLatestVersion loads the latest application version. It will panic if
// called more than once on a running BaseApp.
func (app *BaseApp) LoadLatestVersion(baseKey *sdk.KVStoreKey) error {
	err := app.storeLoader(app.cms)
	if err != nil {
		return err
	}
	return app.initFromMainStore(baseKey)
}

// DefaultStoreLoader will be used by default and loads the latest version
func DefaultStoreLoader(ms sdk.CommitMultiStore) error {
	return ms.LoadLatestVersion()
}

// StoreLoaderWithUpgrade is used to prepare baseapp with a fixed StoreLoader
// pattern. This is useful in test cases, or with custom upgrade loading logic.
func StoreLoaderWithUpgrade(upgrades *storetypes.StoreUpgrades) StoreLoader {
	return func(ms sdk.CommitMultiStore) error {
		return ms.LoadLatestVersionAndUpgrade(upgrades)
	}
}

// UpgradeableStoreLoader can be configured by SetStoreLoader() to check for the
// existence of a given upgrade file - json encoded StoreUpgrades data.
//
// If not file is present, it will peform the default load (no upgrades to store).
//
// If the file is present, it will parse the file and execute those upgrades
// (rename or delete stores), while loading the data. It will also delete the
// upgrade file upon successful load, so that the upgrade is only applied once,
// and not re-applied on next restart
//
// This is useful for in place migrations when a store key is renamed between
// two versions of the software. (TODO: this code will move to x/upgrades
// when PR #4233 is merged, here mainly to help test the design)
func UpgradeableStoreLoader(upgradeInfoPath string) StoreLoader {
	return func(ms sdk.CommitMultiStore) error {
		_, err := os.Stat(upgradeInfoPath)
		if os.IsNotExist(err) {
			return DefaultStoreLoader(ms)
		} else if err != nil {
			return err
		}

		// there is a migration file, let's execute
		data, err := ioutil.ReadFile(upgradeInfoPath)
		if err != nil {
			return fmt.Errorf("Cannot read upgrade file %s: %v", upgradeInfoPath, err)
		}

		var upgrades storetypes.StoreUpgrades
		err = json.Unmarshal(data, &upgrades)
		if err != nil {
			return fmt.Errorf("Cannot parse upgrade file: %v", err)
		}

		err = ms.LoadLatestVersionAndUpgrade(&upgrades)
		if err != nil {
			return fmt.Errorf("Load and upgrade database: %v", err)
		}

		// if we have a successful load, we delete the file
		err = os.Remove(upgradeInfoPath)
		if err != nil {
			return fmt.Errorf("deleting upgrade file %s: %v", upgradeInfoPath, err)
		}
		return nil
	}
}

// LoadVersion loads the BaseApp application version. It will panic if called
// more than once on a running baseapp.
func (app *BaseApp) LoadVersion(version int64, baseKey *sdk.KVStoreKey) error {
	err := app.cms.LoadVersion(version)
	if err != nil {
		return err
	}
	return app.initFromMainStore(baseKey)
}

// LastCommitID returns the last CommitID of the multistore.
func (app *BaseApp) LastCommitID() sdk.CommitID {
	return app.cms.LastCommitID()
}

// LastBlockHeight returns the last committed block height.
func (app *BaseApp) LastBlockHeight() int64 {
	return app.cms.LastCommitID().Version
}

// initializes the remaining logic from app.cms
func (app *BaseApp) initFromMainStore(baseKey *sdk.KVStoreKey) error {
	mainStore := app.cms.GetKVStore(baseKey)
	if mainStore == nil {
		return errors.New("baseapp expects MultiStore with 'main' KVStore")
	}

	// memoize baseKey
	if app.baseKey != nil {
		panic("app.baseKey expected to be nil; duplicate init?")
	}
	app.baseKey = baseKey

	// Load the consensus params from the main store. If the consensus params are
	// nil, it will be saved later during InitChain.
	//
	// TODO: assert that InitChain hasn't yet been called.
	consensusParamsBz := mainStore.Get(mainConsensusParamsKey)
	if consensusParamsBz != nil {
		var consensusParams = &abci.ConsensusParams{}

		err := proto.Unmarshal(consensusParamsBz, consensusParams)
		if err != nil {
			panic(err)
		}

		app.setConsensusParams(consensusParams)
	}

	// needed for the export command which inits from store but never calls initchain
	app.setCheckState(abci.Header{})
	app.Seal()

	return nil
}

func (app *BaseApp) setMinGasPrices(gasPrices sdk.DecCoins) {
	app.minGasPrices = gasPrices
}

func (app *BaseApp) setHaltHeight(height uint64) {
	app.haltHeight = height
}

// Router returns the router of the BaseApp.
func (app *BaseApp) Router() sdk.Router {
	if app.sealed {
		// We cannot return a router when the app is sealed because we can't have
		// any routes modified which would cause unexpected routing behavior.
		panic("Router() on sealed BaseApp")
	}
	return app.router
}

// QueryRouter returns the QueryRouter of a BaseApp.
func (app *BaseApp) QueryRouter() sdk.QueryRouter { return app.queryRouter }

// Seal seals a BaseApp. It prohibits any further modifications to a BaseApp.
func (app *BaseApp) Seal() { app.sealed = true }

// IsSealed returns true if the BaseApp is sealed and false otherwise.
func (app *BaseApp) IsSealed() bool { return app.sealed }

// setCheckState sets checkState with the cached multistore and
// the context wrapping it.
// It is called by InitChain() and Commit()
func (app *BaseApp) setCheckState(header abci.Header) {
	ms := app.cms.CacheMultiStore()
	app.checkState = &state{
		ms:  ms,
		ctx: sdk.NewContext(ms, header, true, app.logger).WithMinGasPrices(app.minGasPrices),
	}
}

// setCheckState sets checkState with the cached multistore and
// the context wrapping it.
// It is called by InitChain() and BeginBlock(),
// and deliverState is set nil on Commit().
func (app *BaseApp) setDeliverState(header abci.Header) {
	ms := app.cms.CacheMultiStore()
	app.deliverState = &state{
		ms:  ms,
		ctx: sdk.NewContext(ms, header, false, app.logger),
	}
}

// setConsensusParams memoizes the consensus params.
func (app *BaseApp) setConsensusParams(consensusParams *abci.ConsensusParams) {
	app.consensusParams = consensusParams
}

// setConsensusParams stores the consensus params to the main store.
func (app *BaseApp) storeConsensusParams(consensusParams *abci.ConsensusParams) {
	consensusParamsBz, err := proto.Marshal(consensusParams)
	if err != nil {
		panic(err)
	}
	mainStore := app.cms.GetKVStore(app.baseKey)
	mainStore.Set(mainConsensusParamsKey, consensusParamsBz)
}

// getMaximumBlockGas gets the maximum gas from the consensus params. It panics
// if maximum block gas is less than negative one and returns zero if negative
// one.
func (app *BaseApp) getMaximumBlockGas() uint64 {
	if app.consensusParams == nil || app.consensusParams.Block == nil {
		return 0
	}

	maxGas := app.consensusParams.Block.MaxGas
	switch {
	case maxGas < -1:
		panic(fmt.Sprintf("invalid maximum block gas: %d", maxGas))

	case maxGas == -1:
		return 0

	default:
		return uint64(maxGas)
	}
}

func (app *BaseApp) validateHeight(req abci.RequestBeginBlock) error {
	if req.Header.Height < 1 {
		return fmt.Errorf("invalid height: %d", req.Header.Height)
	}

	prevHeight := app.LastBlockHeight()
	if req.Header.Height != prevHeight+1 {
		return fmt.Errorf("invalid height: %d; expected: %d", req.Header.Height, prevHeight+1)
	}

	return nil
}

// validateBasicTxMsgs executes basic validator calls for messages.
func validateBasicTxMsgs(msgs []sdk.Msg) sdk.Error {
	if msgs == nil || len(msgs) == 0 {
		return sdk.ErrUnknownRequest("Tx.GetMsgs() must return at least one message in list")
	}

	for _, msg := range msgs {
		// Validate the Msg.
		err := msg.ValidateBasic()
		if err != nil {
			return err
		}
	}

	return nil
}

// retrieve the context for the tx w/ txBytes and other memoized values.
func (app *BaseApp) getContextForTx(mode runTxMode, txBytes []byte) (ctx sdk.Context) {
	ctx = app.getState(mode).ctx.
		WithTxBytes(txBytes).
		WithVoteInfos(app.voteInfos).
		WithConsensusParams(app.consensusParams)

	if mode == runTxModeSimulate {
		ctx, _ = ctx.CacheContext()
	}

	return
}

// runMsgs iterates through all the messages and executes them.
// nolint: gocyclo
func (app *BaseApp) runMsgs(ctx sdk.Context, msgs []sdk.Msg, mode runTxMode) (result sdk.Result) {
	idxLogs := make([]sdk.ABCIMessageLog, 0, len(msgs)) // a list of JSON-encoded logs with msg index

	var (
		data      []byte
		code      sdk.CodeType
		codespace sdk.CodespaceType
	)

	events := sdk.EmptyEvents()

	// NOTE: GasWanted is determined by ante handler and GasUsed by the GasMeter.
	for msgIdx, msg := range msgs {
		// match message route
		msgRoute := msg.Route()
		handler := app.router.Route(msgRoute)
		if handler == nil {
			return sdk.ErrUnknownRequest("Unrecognized Msg type: " + msgRoute).Result()
		}

		var msgResult sdk.Result

		// skip actual execution for CheckTx mode
		if mode != runTxModeCheck {
			msgResult = handler(ctx, msg)
		}

		// Each message result's Data must be length prefixed in order to separate
		// each result.
		data = append(data, msgResult.Data...)

		// append events from the message's execution and a message action event
		events = events.AppendEvent(sdk.NewEvent(sdk.EventTypeMessage, sdk.NewAttribute(sdk.AttributeKeyAction, msg.Type())))
		events = events.AppendEvents(msgResult.Events)

		idxLog := sdk.ABCIMessageLog{MsgIndex: uint16(msgIdx), Log: msgResult.Log}

		// stop execution and return on first failed message
		if !msgResult.IsOK() {
			idxLog.Success = false
			idxLogs = append(idxLogs, idxLog)

			code = msgResult.Code
			codespace = msgResult.Codespace
			break
		}

		idxLog.Success = true
		idxLogs = append(idxLogs, idxLog)
	}

	logJSON := codec.Cdc.MustMarshalJSON(idxLogs)
	result = sdk.Result{
		Code:      code,
		Codespace: codespace,
		Data:      data,
		Log:       strings.TrimSpace(string(logJSON)),
		GasUsed:   ctx.GasMeter().GasConsumed(),
		Events:    events,
	}

	return result
}

// Returns the applications's deliverState if app is in runTxModeDeliver,
// otherwise it returns the application's checkstate.
func (app *BaseApp) getState(mode runTxMode) *state {
	if mode == runTxModeCheck || mode == runTxModeSimulate {
		return app.checkState
	}

	return app.deliverState
}

// cacheTxContext returns a new context based off of the provided context with
// a cache wrapped multi-store.
func (app *BaseApp) cacheTxContext(ctx sdk.Context, txBytes []byte) (
	sdk.Context, sdk.CacheMultiStore) {

	ms := ctx.MultiStore()
	// TODO: https://github.com/cosmos/cosmos-sdk/issues/2824
	msCache := ms.CacheMultiStore()
	if msCache.TracingEnabled() {
		msCache = msCache.SetTracingContext(
			sdk.TraceContext(
				map[string]interface{}{
					"txHash": fmt.Sprintf("%X", tmhash.Sum(txBytes)),
				},
			),
		).(sdk.CacheMultiStore)
	}

	return ctx.WithMultiStore(msCache), msCache
}

// runTx processes a transaction. The transactions is processed via an
// anteHandler. The provided txBytes may be nil in some cases, eg. in tests. For
// further details on transaction execution, reference the BaseApp SDK
// documentation.
func (app *BaseApp) runTx(mode runTxMode, txBytes []byte, tx sdk.Tx) (result sdk.Result) {
	// NOTE: GasWanted should be returned by the AnteHandler. GasUsed is
	// determined by the GasMeter. We need access to the context to get the gas
	// meter so we initialize upfront.
	var gasWanted uint64

	ctx := app.getContextForTx(mode, txBytes)
	ms := ctx.MultiStore()

	// only run the tx if there is block gas remaining
	if mode == runTxModeDeliver && ctx.BlockGasMeter().IsOutOfGas() {
		return sdk.ErrOutOfGas("no block gas left to run tx").Result()
	}

	var startingGas uint64
	if mode == runTxModeDeliver {
		startingGas = ctx.BlockGasMeter().GasConsumed()
	}

	defer func() {
		if r := recover(); r != nil {
			switch rType := r.(type) {
			case sdk.ErrorOutOfGas:
				log := fmt.Sprintf(
					"out of gas in location: %v; gasWanted: %d, gasUsed: %d",
					rType.Descriptor, gasWanted, ctx.GasMeter().GasConsumed(),
				)
				result = sdk.ErrOutOfGas(log).Result()
			default:
				log := fmt.Sprintf("recovered: %v\nstack:\n%v", r, string(debug.Stack()))
				result = sdk.ErrInternal(log).Result()
			}
		}

		result.GasWanted = gasWanted
		result.GasUsed = ctx.GasMeter().GasConsumed()
	}()

	// If BlockGasMeter() panics it will be caught by the above recover and will
	// return an error - in any case BlockGasMeter will consume gas past the limit.
	//
	// NOTE: This must exist in a separate defer function for the above recovery
	// to recover from this one.
	defer func() {
		if mode == runTxModeDeliver {
			ctx.BlockGasMeter().ConsumeGas(
				ctx.GasMeter().GasConsumedToLimit(),
				"block gas meter",
			)

			if ctx.BlockGasMeter().GasConsumed() < startingGas {
				panic(sdk.ErrorGasOverflow{Descriptor: "tx gas summation"})
			}
		}
	}()

	var msgs = tx.GetMsgs()
	if err := validateBasicTxMsgs(msgs); err != nil {
		return err.Result()
	}

	if app.anteHandler != nil {
		var anteCtx sdk.Context
		var msCache sdk.CacheMultiStore

		// Cache wrap context before anteHandler call in case it aborts.
		// This is required for both CheckTx and DeliverTx.
		// Ref: https://github.com/cosmos/cosmos-sdk/issues/2772
		//
		// NOTE: Alternatively, we could require that anteHandler ensures that
		// writes do not happen if aborted/failed.  This may have some
		// performance benefits, but it'll be more difficult to get right.
		anteCtx, msCache = app.cacheTxContext(ctx, txBytes)

		newCtx, result, abort := app.anteHandler(anteCtx, tx, mode == runTxModeSimulate)
		if !newCtx.IsZero() {
			// At this point, newCtx.MultiStore() is cache-wrapped, or something else
			// replaced by the ante handler. We want the original multistore, not one
			// which was cache-wrapped for the ante handler.
			//
			// Also, in the case of the tx aborting, we need to track gas consumed via
			// the instantiated gas meter in the ante handler, so we update the context
			// prior to returning.
			ctx = newCtx.WithMultiStore(ms)
		}

		gasWanted = result.GasWanted

		if abort {
			return result
		}

		msCache.Write()
	}

	// Create a new context based off of the existing context with a cache wrapped
	// multi-store in case message processing fails.
	runMsgCtx, msCache := app.cacheTxContext(ctx, txBytes)
	result = app.runMsgs(runMsgCtx, msgs, mode)
	result.GasWanted = gasWanted

	// Safety check: don't write the cache state unless we're in DeliverTx.
	if mode != runTxModeDeliver {
		return result
	}

	// only update state if all messages pass
	if result.IsOK() {
		msCache.Write()
	}

	return result
}

// Splits a string path using the delimiter '/'.
// e.g. "this/is/funny" becomes []string{"this", "is", "funny"}
func splitPath(requestPath string) (path []string) {
	path = strings.Split(requestPath, "/")
	// first element is empty string
	if len(path) > 0 && path[0] == "" {
		path = path[1:]
	}
	return path
}
