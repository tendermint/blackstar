package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/iavl"
	db "github.com/tendermint/tmlibs/db"
)

type keyVal struct {
	key string
	val string
}

func (kv keyVal) getKV() ([]byte, []byte) {
	return []byte(kv.key), []byte(kv.val)
}

type round []keyVal

// make sure the commits are deterministic
func TestStateCommitHash(t *testing.T) {
	assert, require := assert.New(t), require.New(t)

	cases := [...]struct {
		rounds []round
	}{
		// simple, two rounds, no overlap
		0: {
			[]round{
				[]keyVal{{"abc", "123"}, {"def", "456"}},
				[]keyVal{{"more", "news"}, {"good", "news"}},
			},
		},
		// more complex, order should change if applyCache is not deterministic
		1: {
			[]round{
				[]keyVal{{"abc", "123"}, {"def", "456"}, {"foo", "789"}, {"dings", "646"}},
				[]keyVal{{"hf", "123"}, {"giug", "456"}, {"kgiuvgi", "789"}, {"kjguvgk", "646"}},
				[]keyVal{{"one", "more"}, {"two", "things"}, {"uh", "oh"}, {"a", "2"}},
			},
		},
		// make sure ordering works with overwriting as well
		2: {
			[]round{
				[]keyVal{{"abc", "123"}, {"def", "456"}, {"foo", "789"}, {"dings", "646"}},
				[]keyVal{{"hf", "123"}, {"giug", "456"}, {"kgiuvgi", "789"}, {"kjguvgk", "646"}},
				[]keyVal{{"abc", "qqq"}, {"def", "www"}, {"foo", "ee"}, {"dings", "ff"}},
				[]keyVal{{"one", "more"}, {"uh", "oh"}, {"a", "2"}},
				[]keyVal{{"hf", "dd"}, {"giug", "gg"}, {"kgiuvgi", "jj"}, {"kjguvgk", "uu"}},
			},
		},
	}

	for i, tc := range cases {
		// let's run all rounds... they must each be different,
		// and they must have the same results each run
		var hashes [][]byte

		// try each 5 times for deterministic check
		for j := 0; j < 5; j++ {
			result := make([][]byte, len(tc.rounds))

			// make the store...
			tree := iavl.NewVersionedTree(db.NewMemDB(), 0)
			store := NewState(tree, 2)

			for n, r := range tc.rounds {
				// start the cache
				deliver := store.Append()
				for _, kv := range r {
					// add the value to cache
					k, v := kv.getKV()
					deliver.Set(k, v)
				}
				// commit and add hash to result
				hash, err := store.Commit(int64(n + 1))
				require.Nil(err, "tc:%d / rnd:%d - %+v", i, n, err)
				result[n] = hash
			}

			// make sure result is all unique
			for n := 0; n < len(result)-1; n++ {
				assert.NotEqual(result[n], result[n+1], "tc:%d / rnd:%d", i, n)
			}

			// if hashes != nil, make sure same as last trial
			if hashes != nil {
				for n := 0; n < len(result); n++ {
					assert.Equal(hashes[n], result[n], "tc:%d / rnd:%d", i, n)
				}
			}
			// store to check against next trial
			hashes = result
		}
	}

}

// Ensure getting data doesn't cause it to write
func TestGetDoesntWrite(t *testing.T) {
	assert, require := assert.New(t), require.New(t)

	// make the store...
	tree := iavl.NewVersionedTree(db.NewMemDB(), 0)
	store := NewState(tree, 5)

	k, v := []byte("foo"), []byte("bar")
	k2, v2 := []byte("abc"), []byte("def")
	nok := []byte("baz")

	// one set should change something
	store.Append().Set(k, v)
	hash, err := store.Commit(1)
	require.NoError(err)
	require.NotNil(hash)
	// calling hash returns last committed state
	hash1 := store.LatestHash()
	require.Equal(hash, hash1)

	// a second set will update the state
	store.Append().Set(k2, v2)
	hash, err = store.Commit(2)
	require.NoError(err)
	assert.NotEqual(hash1, hash)
	// calling hash returns last committed state
	hash2 := store.LatestHash()
	require.Equal(hash, hash2)

	// a missed get will not do anything....
	val := store.Append().Get(nok)
	assert.Nil(val)
	hash, err = store.Commit(3)
	require.NoError(err)
	assert.Equal(hash2, hash)

	// a proper get will not do anything....
	val = store.Append().Get(k2)
	assert.Equal(v2, val)
	hash, err = store.Commit(4)
	require.NoError(err)
	assert.Equal(hash2, hash)
}
