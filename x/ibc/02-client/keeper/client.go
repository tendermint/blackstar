package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
	"github.com/cosmos/cosmos-sdk/x/ibc/02-client/types"
	tendermint "github.com/cosmos/cosmos-sdk/x/ibc/07-tendermint"
	ibctmtypes "github.com/cosmos/cosmos-sdk/x/ibc/07-tendermint/types"
)

// CreateClient creates a new client state and populates it with a given consensus
// state as defined in https://github.com/cosmos/ics/tree/master/spec/ics-002-client-semantics#create
//
// CONTRACT: ClientState was constructed correctly from given initial consensusState
func (k Keeper) CreateClient(
	ctx sdk.Context, clientState exported.ClientState, consensusState exported.ConsensusState,
) (exported.ClientState, error) {
	clientID := clientState.GetID()
	_, found := k.GetClientState(ctx, clientID)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrClientExists, "cannot create client with ID %s", clientID)
	}

	_, found = k.GetClientType(ctx, clientID)
	if found {
		panic(fmt.Sprintf("client type is already defined for client %s", clientID))
	}

	height := consensusState.GetHeight()
	if consensusState != nil {
		k.SetClientConsensusState(ctx, clientID, height, consensusState)
	}

	k.SetClientState(ctx, clientState)
	k.SetClientType(ctx, clientID, clientState.ClientType())
	k.Logger(ctx).Info(fmt.Sprintf("client %s created at height %d", clientID, clientState.GetLatestHeight()))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateClient,
			sdk.NewAttribute(types.AttributeKeyClientID, clientID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return clientState, nil
}

// UpdateClient updates the consensus state and the state root from a provided header
func (k Keeper) UpdateClient(ctx sdk.Context, clientID string, header exported.Header) error {
	clientType, found := k.GetClientType(ctx, clientID)
	if !found {
		return sdkerrors.Wrapf(types.ErrClientTypeNotFound, "cannot update client with ID %s", clientID)
	}

	// check that the header consensus matches the client one
	if header.ClientType() != clientType {
		return sdkerrors.Wrapf(types.ErrInvalidConsensus, "cannot update client with ID %s", clientID)
	}

	clientState, found := k.GetClientState(ctx, clientID)
	if !found {
		return sdkerrors.Wrapf(types.ErrClientNotFound, "cannot update client with ID %s", clientID)
	}

	// addittion to spec: prevent update if the client is frozen
	if clientState.IsFrozen() {
		return sdkerrors.Wrapf(types.ErrClientFrozen, "cannot update client with ID %s", clientID)
	}

	var (
		consensusState exported.ConsensusState
		err            error
	)

	switch clientType {
	case exported.Tendermint:
		clientState, consensusState, err = tendermint.CheckValidityAndUpdateState(
			clientState, header, ctx.BlockTime(),
		)
	default:
		err = types.ErrInvalidClientType
	}

	if err != nil {
		return sdkerrors.Wrapf(err, "cannot update client with ID %s", clientID)
	}

	k.SetClientState(ctx, clientState)
	fmt.Println(clientID)
	fmt.Println(header.GetHeight())
	k.SetClientConsensusState(ctx, clientID, header.GetHeight(), consensusState)
	k.Logger(ctx).Info(fmt.Sprintf("client %s updated to height %d", clientID, header.GetHeight()))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateClient,
			sdk.NewAttribute(types.AttributeKeyClientID, clientID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return nil
}

// CheckMisbehaviourAndUpdateState checks for client misbehaviour and freezes the
// client if so.
func (k Keeper) CheckMisbehaviourAndUpdateState(ctx sdk.Context, misbehaviour exported.Misbehaviour) error {
	clientState, found := k.GetClientState(ctx, misbehaviour.GetClientID())
	if !found {
		return sdkerrors.Wrap(types.ErrClientNotFound, misbehaviour.GetClientID())
	}

	consensusState, found := k.GetClientConsensusStateLTE(ctx, misbehaviour.GetClientID(), uint64(misbehaviour.GetHeight()))
	if !found {
		return sdkerrors.Wrap(types.ErrConsensusStateNotFound, misbehaviour.GetClientID())
	}

	var err error
	switch e := misbehaviour.(type) {
	case ibctmtypes.Evidence:
		clientState, err = tendermint.CheckMisbehaviourAndUpdateState(
			clientState, consensusState, misbehaviour, consensusState.GetHeight(), ctx.BlockTime(),
		)

	default:
		err = sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized IBC client evidence type: %T", e)
	}

	if err != nil {
		return err
	}

	k.SetClientState(ctx, clientState)
	k.Logger(ctx).Info(fmt.Sprintf("client %s frozen due to misbehaviour", misbehaviour.GetClientID()))

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSubmitMisbehaviour,
			sdk.NewAttribute(types.AttributeKeyClientID, misbehaviour.GetClientID()),
			sdk.NewAttribute(types.AttributeKeyClientType, misbehaviour.ClientType().String()),
		),
	)

	return nil
}
