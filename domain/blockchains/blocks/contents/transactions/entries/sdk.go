package entries

import "github.com/steve-care-software/steve/domain/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the entry adapter
type Adapter interface {
	ToBytes(ins Entry) ([]byte, error)
	ToInstance(data []byte) (Entry, error)
}

// Builder represents the entry builder
type Builder interface {
	Create() Builder
	WithFlag(flag hash.Hash) Builder
	WithScript(script hash.Hash) Builder
	WithFees(fees uint64) Builder
	Now() (Entry, error)
}

// Entry represents an entry
type Entry interface {
	Hash() hash.Hash
	Flag() hash.Hash
	Script() hash.Hash
	Fees() uint64
}
