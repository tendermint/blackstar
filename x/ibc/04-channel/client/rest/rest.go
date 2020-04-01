package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/ibc/04-channel/exported"
	"github.com/cosmos/cosmos-sdk/x/ibc/04-channel/types"
	commitmentexported "github.com/cosmos/cosmos-sdk/x/ibc/23-commitment/exported"
)

const (
	RestChannelID = "channel-id"
	RestPortID    = "port-id"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, queryRoute string) {
	registerQueryRoutes(cliCtx, r, queryRoute)
	registerTxRoutes(cliCtx, r)
}

// ChannelOpenInitReq defines the properties of a channel open init request's body.
type ChannelOpenInitReq struct {
	BaseReq               rest.BaseReq   `json:"base_req" yaml:"base_req"`
	PortID                string         `json:"port_id" yaml:"port_id"`
	ChannelID             string         `json:"channel_id" yaml:"channel_id"`
	Version               string         `json:"version" yaml:"version"`
	ChannelOrder          exported.Order `json:"channel_order" yaml:"channel_order"`
	ConnectionHops        []string       `json:"connection_hops" yaml:"connection_hops"`
	CounterpartyPortID    string         `json:"counterparty_port_id" yaml:"counterparty_port_id"`
	CounterpartyChannelID string         `json:"counterparty_channel_id" yaml:"counterparty_channel_id"`
}

// ChannelTryOpenReq defines the properties of a channel open try request's body.
type ChannelTryOpenReq struct {
	BaseReq               rest.BaseReq             `json:"base_req" yaml:"base_req"`
	PortID                string                   `json:"port_id" yaml:"port_id"`
	ChannelID             string                   `json:"channel_id" yaml:"channel_id"`
	Version               string                   `json:"version" yaml:"version"`
	ChannelOrder          exported.Order           `json:"channel_order" yaml:"channel_order"`
	ConnectionHops        []string                 `json:"connection_hops" yaml:"connection_hops"`
	CounterpartyPortID    string                   `json:"counterparty_port_id" yaml:"counterparty_port_id"`
	CounterpartyChannelID string                   `json:"counterparty_channel_id" yaml:"counterparty_channel_id"`
	CounterpartyVersion   string                   `json:"counterparty_version" yaml:"counterparty_version"`
	ProofInit             commitmentexported.Proof `json:"proof_init" yaml:"proof_init"`
	ProofHeight           uint64                   `json:"proof_height" yaml:"proof_height"`
}

// ChannelOpenAckReq defines the properties of a channel open ack request's body.
type ChannelOpenAckReq struct {
	BaseReq             rest.BaseReq             `json:"base_req" yaml:"base_req"`
	CounterpartyVersion string                   `json:"counterparty_version" yaml:"counterparty_version"`
	ProofTry            commitmentexported.Proof `json:"proof_try" yaml:"proof_try"`
	ProofHeight         uint64                   `json:"proof_height" yaml:"proof_height"`
}

// ChannelOpenConfirmReq defines the properties of a channel open confirm request's body.
type ChannelOpenConfirmReq struct {
	BaseReq     rest.BaseReq             `json:"base_req" yaml:"base_req"`
	ProofAck    commitmentexported.Proof `json:"proof_ack" yaml:"proof_ack"`
	ProofHeight uint64                   `json:"proof_height" yaml:"proof_height"`
}

// ConnectionOpenInitReq defines the properties of a channel close init request's body.
type ChannelCloseInitReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
}

// ChannelCloseConfirmReq defines the properties of a channel close confirm request's body.
type ChannelCloseConfirmReq struct {
	BaseReq     rest.BaseReq             `json:"base_req" yaml:"base_req"`
	ProofInit   commitmentexported.Proof `json:"proof_init" yaml:"proof_init"`
	ProofHeight uint64                   `json:"proof_height" yaml:"proof_height"`
}

// RecvPacketReq defines the properties of a receive packet request's body.
type RecvPacketReq struct {
	BaseReq rest.BaseReq             `json:"base_req" yaml:"base_req"`
	Packet  types.Packet             `json:"packet" yaml:"packet"`
	Proofs  commitmentexported.Proof `json:"proofs" yaml:"proofs"`
	Height  uint64                   `json:"height" yaml:"height"`
}
