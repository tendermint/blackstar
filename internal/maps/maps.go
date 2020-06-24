package maps

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/tendermint/tendermint/crypto/merkle"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/tendermint/tendermint/libs/kv"
)

// MerkleMap defines a merkle-ized tree from a map. Leave values are treated as
// hash(key) | hash(value). Leaves are sorted before Merkle hashing.
type MerkleMap struct {
	kvs    kv.Pairs
	sorted bool
}

func NewMerkleMap() *MerkleMap {
	return &MerkleMap{
		kvs:    nil,
		sorted: false,
	}
}

// Set creates a kv.Pair from the provided key and value. The value is hashed prior
// to creating a kv.Pair. The created kv.Pair is appended to the MerkleMap's slice
// of kv.Pairs. Whenever called, the MerkleMap must be resorted.
func (sm *MerkleMap) Set(key string, value []byte) {
	sm.sorted = false

	// The value is hashed, so you can check for equality with a cached value (say)
	// and make a determination to fetch or not.
	vhash := tmhash.Sum(value)

	sm.kvs = append(sm.kvs, kv.Pair{
		Key:   []byte(key),
		Value: vhash,
	})
}

// Hash returns the merkle root of items sorted by key. Note, it is unstable.
func (sm *MerkleMap) Hash() []byte {
	sm.sort()
	return hashKVPairs(sm.kvs)
}

func (sm *MerkleMap) sort() {
	if sm.sorted {
		return
	}

	sm.kvs.Sort()
	sm.sorted = true
}

// kvPair defines a type alias for kv.Pair so that we can create bytes to hash
// when constructing the merkle root. Note, key and values are both length-prefixed.
type kvPair kv.Pair

// bytes returns a byte slice representation of the kvPair where the key and value
// are length-prefixed.
func (kv kvPair) bytes() []byte {
	var b bytes.Buffer

	err := encodeByteSlice(&b, kv.Key)
	if err != nil {
		panic(err)
	}

	err = encodeByteSlice(&b, kv.Value)
	if err != nil {
		panic(err)
	}

	return b.Bytes()
}

func encodeByteSlice(w io.Writer, bz []byte) error {
	var buf [8]byte
	n := binary.PutUvarint(buf[:], uint64(len(bz)))

	_, err := w.Write(buf[:n])
	if err != nil {
		return err
	}

	_, err = w.Write(bz)
	return err
}

// hashKVPairs hashes a kvPair and creates a merkle tree where the leaves are
// byte slices.
func hashKVPairs(kvs kv.Pairs) []byte {
	kvsH := make([][]byte, len(kvs))
	for i, kvp := range kvs {
		kvsH[i] = kvPair(kvp).bytes()
	}

	return merkle.SimpleHashFromByteSlices(kvsH)
}

// ---------------------------------------------

// Merkle tree from a map.
// Leaves are `hash(key) | hash(value)`.
// Leaves are sorted before Merkle hashing.
type SimpleMap struct {
	Kvs    kv.Pairs
	sorted bool
}

func NewSimpleMap() *SimpleMap {
	return &SimpleMap{
		Kvs:    nil,
		sorted: false,
	}
}

// Set creates a kv pair of the key and the hash of the value,
// and then appends it to SimpleMap's kv pairs.
func (sm *SimpleMap) Set(key string, value []byte) {
	sm.sorted = false

	// The value is hashed, so you can
	// check for equality with a cached value (say)
	// and make a determination to fetch or not.
	vhash := tmhash.Sum(value)

	sm.Kvs = append(sm.Kvs, kv.Pair{
		Key:   []byte(key),
		Value: vhash,
	})
}

// Hash Merkle root hash of items sorted by key
// (UNSTABLE: and by value too if duplicate key).
func (sm *SimpleMap) Hash() []byte {
	sm.Sort()
	return hashKVPairs(sm.Kvs)
}

func (sm *SimpleMap) Sort() {
	if sm.sorted {
		return
	}
	sm.Kvs.Sort()
	sm.sorted = true
}

// Returns a copy of sorted KVPairs.
// NOTE these contain the hashed key and value.
func (sm *SimpleMap) KVPairs() kv.Pairs {
	sm.Sort()
	kvs := make(kv.Pairs, len(sm.Kvs))
	copy(kvs, sm.Kvs)
	return kvs
}

//----------------------------------------

// A local extension to KVPair that can be hashed.
// Key and value are length prefixed and concatenated,
// then hashed.
type KVPair kv.Pair

// NewKVPair takes in a key and value and creates a kv.Pair
// wrapped in the local extension KVPair
func NewKVPair(key, value []byte) KVPair {
	return KVPair(kv.Pair{
		Key:   key,
		Value: value,
	})
}

// Bytes returns key || value, with both the
// key and value length prefixed.
func (kv KVPair) Bytes() []byte {
	var b bytes.Buffer
	err := encodeByteSlice(&b, kv.Key)
	if err != nil {
		panic(err)
	}
	err = encodeByteSlice(&b, kv.Value)
	if err != nil {
		panic(err)
	}
	return b.Bytes()
}

// SimpleHashFromMap computes a merkle tree from sorted map and returns the merkle
// root.
func SimpleHashFromMap(m map[string][]byte) []byte {
	mm := NewMerkleMap()
	for k, v := range m {
		mm.Set(k, v)
	}

	return mm.Hash()
}

// SimpleProofsFromMap generates proofs from a map. The keys/values of the map will be used as the keys/values
// in the underlying key-value pairs.
// The keys are sorted before the proofs are computed.
func SimpleProofsFromMap(m map[string][]byte) (rootHash []byte, proofs map[string]*merkle.SimpleProof, keys []string) {
	sm := NewSimpleMap()
	for k, v := range m {
		sm.Set(k, v)
	}
	sm.Sort()
	kvs := sm.Kvs
	kvsBytes := make([][]byte, len(kvs))
	for i, kvp := range kvs {
		kvsBytes[i] = KVPair(kvp).Bytes()
	}

	rootHash, proofList := merkle.SimpleProofsFromByteSlices(kvsBytes)
	proofs = make(map[string]*merkle.SimpleProof)
	keys = make([]string, len(proofList))
	for i, kvp := range kvs {
		proofs[string(kvp.Key)] = proofList[i]
		keys[i] = string(kvp.Key)
	}
	return
}
