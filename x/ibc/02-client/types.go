package client

import (
	"github.com/cosmos/cosmos-sdk/x/ibc/23-commitment"
)

type ValidityPredicateBase interface {
	Kind() Kind
	GetHeight() int64
	Equal(ValidityPredicateBase) bool
}

// ConsensusState
type Client interface {
	Kind() Kind
	GetBase() ValidityPredicateBase
	GetRoot() commitment.Root
	Validate(Header) (Client, error) // ValidityPredicate
}

func Equal(client1, client2 Client) bool {
	return client1.Kind() == client2.Kind() &&
		client1.GetBase().Equal(client2.GetBase())
}

type Header interface {
	Kind() Kind
	//	Proof() HeaderProof
	Base() ValidityPredicateBase // can be nil
	GetRoot() commitment.Root
}

// XXX: Kind should be enum?

type Kind byte

const (
	Tendermint Kind = iota
)
