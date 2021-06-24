package dbtest

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	dbm "github.com/cosmos/cosmos-sdk/db"
)

const strChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" // 62 characters

// RandStr constructs a random alphanumeric string of given length.
func RandStr(length int) string {
	chars := []byte{}
MAIN_LOOP:
	for {
		val := rand.Int63() // nolint:gosec // G404: Use of weak random number generator
		for i := 0; i < 10; i++ {
			v := int(val & 0x3f) // rightmost 6 bits
			if v >= 62 {         // only 62 characters in strChars
				val >>= 6
				continue
			} else {
				chars = append(chars, strChars[v])
				if len(chars) == length {
					break MAIN_LOOP
				}
				val >>= 6
			}
		}
	}

	return string(chars)
}

func Int642Bytes(i int64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func Bytes2Int64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func BenchmarkRangeScans(b *testing.B, db dbm.DBReadWriter, dbSize int64) {
	b.StopTimer()

	rangeSize := int64(10000)
	if dbSize < rangeSize {
		b.Errorf("db size %v cannot be less than range size %v", dbSize, rangeSize)
	}

	for i := int64(0); i < dbSize; i++ {
		bytes := Int642Bytes(i)
		err := db.Set(bytes, bytes)
		if err != nil {
			// require.NoError() is very expensive (according to profiler), so check manually
			b.Fatal(b, err)
		}
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {

		start := rand.Int63n(dbSize - rangeSize) // nolint: gosec
		end := start + rangeSize
		iter, err := db.Iterator(Int642Bytes(start), Int642Bytes(end))
		require.NoError(b, err)
		count := 0
		for ; iter.Valid(); iter.Next() {
			count++
		}
		iter.Close()
		require.EqualValues(b, rangeSize, count)
	}
}

func BenchmarkRandomReadsWrites(b *testing.B, db dbm.DBReadWriter) {
	b.StopTimer()

	// create dummy data
	const numItems = int64(1000000)
	internal := map[int64]int64{}
	for i := 0; i < int(numItems); i++ {
		internal[int64(i)] = int64(0)
	}

	// fmt.Println("ok, starting")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		// Write something
		{
			idx := rand.Int63n(numItems) // nolint: gosec
			internal[idx]++
			val := internal[idx]
			idxBytes := Int642Bytes(idx)
			valBytes := Int642Bytes(val)
			// fmt.Printf("Set %X -> %X\n", idxBytes, valBytes)
			err := db.Set(idxBytes, valBytes)
			if err != nil {
				// require.NoError() is very expensive (according to profiler), so check manually
				b.Fatal(b, err)
			}
		}

		// Read something
		{
			idx := rand.Int63n(numItems) // nolint: gosec
			valExp := internal[idx]
			idxBytes := Int642Bytes(idx)
			valBytes, err := db.Get(idxBytes)
			if err != nil {
				// require.NoError() is very expensive (according to profiler), so check manually
				b.Fatal(b, err)
			}
			// fmt.Printf("Get %X -> %X\n", idxBytes, valBytes)
			if valExp == 0 {
				if !bytes.Equal(valBytes, nil) {
					b.Errorf("Expected %v for %v, got %X", nil, idx, valBytes)
					break
				}
			} else {
				if len(valBytes) != 8 {
					b.Errorf("Expected length 8 for %v, got %X", idx, valBytes)
					break
				}
				valGot := Bytes2Int64(valBytes)
				if valExp != valGot {
					b.Errorf("Expected %v for %v, got %v", valExp, idx, valGot)
					break
				}
			}
		}

	}
}
