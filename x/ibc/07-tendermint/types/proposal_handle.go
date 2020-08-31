package types

import (
	"time"

	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clientexported "github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/02-client/types"
)

// CheckProposedHeaderAndUpdateState will try to update the client with the new header if and
// only if the proposal passes and one of the following two conditions is satisfied:
// 		1) AllowGovernanceOverrideAfterExpiry=true and Expire(ctx.BlockTime) = true
// 		2) AllowGovernanceOverrideAfterMinbehaviour and IsFrozen() = true
// In case 2) before trying to update the client, the client will be unfrozen by resetting
// the FrozenHeight to the zero Height. Note, that even if the update happens, it may not be
// successful. The header may fail validation checks and an error will be returned in that
// case.
func (cs ClientState) CheckProposedHeaderAndUpdateState(
	ctx sdk.Context, cdc codec.BinaryMarshaler, clientStore sdk.KVStore,
	header clientexported.Header,
) (clientexported.ClientState, clientexported.ConsensusState, error) {
	tmHeader, ok := header.(*Header)
	if !ok {
		return nil, nil, sdkerrors.Wrapf(
			clienttypes.ErrInvalidHeader, "expected type %T, got %T", &Header{}, header,
		)
	}

	// get consensus state corresponding to client state to check if the client is expired
	consensusState, err := GetConsensusState(clientStore, cdc, cs.GetLatestHeight())
	if err != nil {
		return nil, nil, sdkerrors.Wrapf(
			err, "could not get consensus state from clientstore at height: %d", cs.GetLatestHeight(),
		)
	}

	// unfreeze client if the client is frozen and this is allowed. Otherwise if the client
	// is not expired or not allowed to be updated after expiry then the proposal cannot update
	// the client.
	if cs.IsFrozen() && cs.AllowGovernanceOverrideAfterMisbehaviour {
		cs.FrozenHeight = 0
	} else if !(cs.AllowGovernanceOverrideAfterExpiry && cs.Expired(consensusState.Timestamp, ctx.BlockTime())) {
		return nil, nil, sdkerrors.Wrap(clienttypes.ErrUpdateClientFailed, "client cannot be updated with proposal")
	}

	cs.checkProposedHeader(consensusState, tmHeader, ctx.BlockTime())

	newClientState, consensusState := update(&cs, tmHeader)
	return newClientState, consensusState, nil
}

// checkProposedHeader checks if the Tendermint header is valid for updating a client after
// a passed proposal.
// It returns an error if:
// - the header provided is not parseable to tendermint types
// - header height is less than or equal to the latest client state height
// - signed tendermint header is invalid
// - header timestamp is less than or equal to the latest consensus state timestamp
// - header timestamp is expired
// NOTE: header.ValidateBasic is called in the 02-client proposal handler. Additional checks
// on the validator set and the validator set hash are done in header.ValidateBasic.
func (cs ClientState) checkProposedHeader(consensusState *ConsensusState, header *Header, currentTimestamp time.Time) error {
	tmSignedHeader, err := tmtypes.SignedHeaderFromProto(header.SignedHeader)
	if err != nil {
		return sdkerrors.Wrap(err, "signed header in not tendermint signed header type")
	}

	if !header.GetTime().After(consensusState.Timestamp) {
		return sdkerrors.Wrapf(
			clienttypes.ErrInvalidHeader,
			"header timestamp is less than or equal to latest consensus state timestamp (%s ≤ %s)", header.GetTime(), consensusState.Timestamp)
	}

	// assert header height is newer than latest client state
	if header.GetHeight() <= cs.GetLatestHeight() {
		return sdkerrors.Wrapf(
			clienttypes.ErrInvalidHeader,
			"header height ≤ consensus state height (%d ≤ %d)", header.GetHeight(), cs.GetLatestHeight(),
		)
	}

	if err := tmSignedHeader.ValidateBasic(cs.GetChainID()); err != nil {
		return sdkerrors.Wrap(err, "signed header failed basic validation")
	}

	if cs.Expired(header.GetTime(), currentTimestamp) {
		return sdkerrors.Wrap(clienttypes.ErrInvalidHeader, "header timestamp is already expired")
	}

	return nil
}
