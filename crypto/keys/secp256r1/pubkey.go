package secp256r1

import (
	tmcrypto "github.com/tendermint/tendermint/crypto"

	ecdsa "github.com/cosmos/cosmos-sdk/crypto/keys/internal/ecdsa"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

// String implements proto.Message interface.
func (m *PubKey) String() string {
	return m.Key.String(name)
}

// Bytes implements sdk.PubKey interface.
func (m *PubKey) Bytes() []byte {
	return m.Key.Bytes()
}

// Equals implements sdk.PubKey interface.
func (m *PubKey) Equals(other cryptotypes.PubKey) bool {
	pk2, ok := other.(*PubKey)
	if !ok {
		return false
	}
	return m.Key.Equal(&pk2.Key.PublicKey)
}

// Address implements sdk.PubKey interface.
func (m *PubKey) Address() tmcrypto.Address {
	return m.Key.Address(m.XXX_MessageName())
}

// Type returns key type name. Implements sdk.PubKey interface.
func (m *PubKey) Type() string {
	return name
}

// VerifySignature implements sdk.PubKey interface.
func (m *PubKey) VerifySignature(msg []byte, sig []byte) bool {
	return m.Key.VerifySignature(msg, sig)
}

type ecdsaPK struct {
	ecdsa.PubKey
}

// Size implements proto.Marshaler interface
func (pk *ecdsaPK) Size() int {
	if pk == nil {
		return 0
	}
	return pubKeySize
}

// Unmarshal implements proto.Marshaler interface
func (pk *ecdsaPK) Unmarshal(bz []byte) error {
	return pk.PubKey.Unmarshal(bz, secp256r1, pubKeySize)
}
