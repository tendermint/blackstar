package memdb

import (
	"testing"

	dbm "github.com/cosmos/cosmos-sdk/db"
	"github.com/cosmos/cosmos-sdk/db/dbtest"
)

func BenchmarkMemDBRangeScans1M(b *testing.B) {
	db := NewDB()
	defer db.Close()

	dbtest.BenchmarkRangeScans(b, db, int64(1e6))
}

func BenchmarkMemDBRangeScans10M(b *testing.B) {
	db := NewDB()
	defer db.Close()

	dbtest.BenchmarkRangeScans(b, db, int64(10e6))
}

func BenchmarkMemDBRandomReadsWrites(b *testing.B) {
	db := NewDB()
	defer db.Close()

	dbtest.BenchmarkRandomReadsWrites(b, db)
}

func load(t *testing.T, _ string) dbm.DB {
	return NewDB()
}

func TestGetSetHasDelete(t *testing.T) {
	dbtest.DoTestGetSetHasDelete(t, load)
}

func TestIterators(t *testing.T) {
	dbtest.DoTestIterators(t, load)
}

func TestVersioning(t *testing.T) {
	dbtest.DoTestVersioning(t, load)
}
