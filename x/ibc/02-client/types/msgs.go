package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
	host "github.com/cosmos/cosmos-sdk/x/ibc/24-host"
	ibctypes "github.com/cosmos/cosmos-sdk/x/ibc/types"
)

// Message types for the IBC client
const (
	TypeMsgCreateClient    string = "create_client"
	TypeMsgUpdateClient    string = "update_client"
	TypeClientMisbehaviour string = "client_misbehaviour"
)

var _ sdk.Msg = MsgCreateClient{}

// MsgCreateClient defines a message to create an IBC client
type MsgCreateClient struct {
	ClientID        string                  `json:"client_id" yaml:"client_id"`
	ClientType      string                  `json:"client_type" yaml:"client_type"`
	ConsensusState  exported.ConsensusState `json:"consensus_state" yaml:"consensus_address"`
	TrustingPeriod  time.Duration           `json:"trusting_period" yaml:"trusting_period"`
	UnbondingPeriod time.Duration           `json:"unbonding_period" yaml:"unbonding_period"`
	Signer          sdk.AccAddress          `json:"address" yaml:"address"`
}

// NewMsgCreateClient creates a new MsgCreateClient instance
func NewMsgCreateClient(
	id, clientType string, consensusState exported.ConsensusState,
	trustingPeriod, unbondingPeriod time.Duration, signer sdk.AccAddress,
) MsgCreateClient {
	return MsgCreateClient{
		ClientID:        id,
		ClientType:      clientType,
		ConsensusState:  consensusState,
		TrustingPeriod:  trustingPeriod,
		UnbondingPeriod: unbondingPeriod,
		Signer:          signer,
	}
}

// Route implements sdk.Msg
func (msg MsgCreateClient) Route() string {
	return ibctypes.RouterKey
}

// Type implements sdk.Msg
func (msg MsgCreateClient) Type() string {
	return TypeMsgCreateClient
}

// ValidateBasic implements sdk.Msg
func (msg MsgCreateClient) ValidateBasic() error {
	if clientType := exported.ClientTypeFromString(msg.ClientType); clientType == 0 {
		return sdkerrors.Wrap(ErrInvalidClientType, msg.ClientType)
	}
	if msg.ConsensusState == nil {
		return ErrInvalidConsensus
	}
	if err := msg.ConsensusState.ValidateBasic(); err != nil {
		return err
	}
	if msg.TrustingPeriod == 0 {
		return sdkerrors.Wrap(ErrInvalidTrustingPeriod, "duration cannot be 0")
	}
	if msg.UnbondingPeriod == 0 {
		return sdkerrors.Wrap(ErrInvalidUnbondingPeriod, "duration cannot be 0")
	}
	if msg.Signer.Empty() {
		return sdkerrors.ErrInvalidAddress
	}
	return host.DefaultClientIdentifierValidator(msg.ClientID)
}

// GetSignBytes implements sdk.Msg
func (msg MsgCreateClient) GetSignBytes() []byte {
	return sdk.MustSortJSON(SubModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements sdk.Msg
func (msg MsgCreateClient) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

var _ sdk.Msg = MsgUpdateClient{}

// MsgUpdateClient defines a message to update an IBC client
type MsgUpdateClient struct {
	ClientID  string          `json:"client_id" yaml:"client_id"`
	OldHeader exported.Header `json:"old_header" yaml:"old_header"`
	NewHeader exported.Header `json:"new_header" yaml:"new_header"`
	Signer    sdk.AccAddress  `json:"address" yaml:"address"`
}

// NewMsgUpdateClient creates a new MsgUpdateClient instance
func NewMsgUpdateClient(id string, oldHeader, newHeader exported.Header, signer sdk.AccAddress) MsgUpdateClient {
	return MsgUpdateClient{
		ClientID:  id,
		OldHeader: oldHeader,
		NewHeader: newHeader,
		Signer:    signer,
	}
}

// Route implements sdk.Msg
func (msg MsgUpdateClient) Route() string {
	return ibctypes.RouterKey
}

// Type implements sdk.Msg
func (msg MsgUpdateClient) Type() string {
	return TypeMsgUpdateClient
}

// ValidateBasic implements sdk.Msg
func (msg MsgUpdateClient) ValidateBasic() error {
	if msg.OldHeader == nil {
		return ErrInvalidHeader
	}
	if msg.NewHeader == nil {
		return ErrInvalidHeader
	}
	if msg.Signer.Empty() {
		return sdkerrors.ErrInvalidAddress
	}
	return host.DefaultClientIdentifierValidator(msg.ClientID)
}

// GetSignBytes implements sdk.Msg
func (msg MsgUpdateClient) GetSignBytes() []byte {
	return sdk.MustSortJSON(SubModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements sdk.Msg
func (msg MsgUpdateClient) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}
