package algo

import (
	"github.com/cosmos/go-bip39"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/cosmos/cosmos-sdk/crypto/keys/hd"
)

// PubKeyType defines an algorithm to derive key-pairs which can be used for cryptographic signing.
type PubKeyType string

const (
	// MultiType implies that a pubkey is a multisignature
	MultiType = PubKeyType("multi")
	// Secp256k1Type uses the Bitcoin secp256k1 ECDSA parameters.
	Secp256k1Type = PubKeyType("secp256k1")
	// Ed25519Type represents the Ed25519Type signature system.
	// It is currently not supported for end-user keys (wallets/ledgers).
	Ed25519Type = PubKeyType("ed25519")
	// Sr25519Type represents the Sr25519Type signature system.
	Sr25519Type = PubKeyType("sr25519")
)

var (
	// Secp256k1 uses the Bitcoin secp256k1 ECDSA parameters.
	Secp256k1 = secp256k1Algo{}
)

type DeriveKeyFn func(mnemonic string, bip39Passphrase, hdPath string) ([]byte, error)
type PrivKeyGenFn func(bz []byte) crypto.PrivKey

type secp256k1Algo struct {
}

func (s secp256k1Algo) Name() PubKeyType {
	return Secp256k1Type
}

func (s secp256k1Algo) DeriveKey() DeriveKeyFn {
	return Secp256k1DeriveKey
}

func (s secp256k1Algo) PrivKeyGen() PrivKeyGenFn {
	return Secp256k1PrivKeyGen
}

// Secp256k1PrivKeyGen generates a secp256k1 private key from the given bytes
func Secp256k1PrivKeyGen(bz []byte) crypto.PrivKey {
	var bzArr [32]byte
	copy(bzArr[:], bz)
	return secp256k1.PrivKeySecp256k1(bzArr)
}

// Secp256k1DeriveKey derives and returns the secp256k1 private key for the given seed and HD path.
func Secp256k1DeriveKey(mnemonic string, bip39Passphrase, hdPath string) ([]byte, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, bip39Passphrase)
	if err != nil {
		return nil, err
	}

	masterPriv, ch := hd.ComputeMastersFromSeed(seed)
	if len(hdPath) == 0 {
		return masterPriv[:], nil
	}
	derivedKey, err := hd.DerivePrivateKeyForPath(masterPriv, ch, hdPath)
	return derivedKey[:], err
}
