package entries

import "github.com/steve-care-software/steve/engine/domain/hash"

// NewEntryForTests creates a new entry for tests
func NewEntryForTests(flag hash.Hash, script hash.Hash, fees uint64) Entry {
	ins, err := NewBuilder().Create().
		WithFlag(flag).
		WithScript(script).
		WithFees(fees).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
