package client

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidenceexported "github.com/cosmos/cosmos-sdk/x/evidence/exported"
	"github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
	"github.com/cosmos/cosmos-sdk/x/ibc/02-client/types"
)

// HandleMsgCreateClient defines the sdk.Handler for MsgCreateClient
func HandleMsgCreateClient(ctx sdk.Context, k Keeper, msg MsgCreateClient) (*sdk.Result, error) {
	clientType := exported.ClientTypeFromString(msg.ClientType)
	if clientType == 0 {
		return nil, sdkerrors.Wrap(ErrInvalidClientType, msg.ClientType)
	}

	_, err := k.CreateClient(
		ctx, msg.ClientID, clientType, msg.ConsensusState, msg.TrustingPeriod, msg.UnbondingPeriod,
	)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeCreateClient,
			sdk.NewAttribute(AttributeKeyClientID, msg.ClientID),
			sdk.NewAttribute(AttrbuteKeyClientType, msg.ClientType),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer.String()),
		),
	})

	return &sdk.Result{
		Events: ctx.EventManager().Events(),
	}, nil
}

// HandleMsgUpdateClient defines the sdk.Handler for MsgUpdateClient
func HandleMsgUpdateClient(ctx sdk.Context, k Keeper, msg MsgUpdateClient) (*sdk.Result, error) {
	if err := k.UpdateClient(ctx, msg.ClientID, msg.Header); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeUpdateClient,
			sdk.NewAttribute(AttributeKeyClientID, msg.ClientID),
			sdk.NewAttribute(AttrbuteKeyClientType, msg.Header.ClientType().String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer.String()),
		),
	})

	return &sdk.Result{
		Events: ctx.EventManager().Events(),
	}, nil
}

// HandlerClientMisbehaviour defines the Evidence module handler for submitting a
// light client misbehaviour.
func HandlerClientMisbehaviour(k Keeper) evidence.Handler {
	return func(ctx sdk.Context, evidence evidenceexported.Evidence) error {
		misbehaviour, ok := evidence.(exported.Misbehaviour)
		if !ok {
			return types.ErrInvalidEvidence
		}

		return k.CheckMisbehaviourAndUpdateState(ctx, misbehaviour)
	}
}
