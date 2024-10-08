package entries

import "github.com/steve-care-software/steve/hash"

const dataLengthTooSmallErrPattern = "(entry) the data length was expected to be at least %d bytes, %d returned"

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	hashAdapter := hash.NewAdapter()
	builder := NewBuilder()
	return createAdapter(
		hashAdapter,
		builder,
	)
}

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
	ToInstance(data []byte) (Entry, []byte, error)
}

// Builder represents the entry builder
type Builder interface {
	Create() Builder
	WithFlag(flag hash.Hash) Builder
	WithScript(script []byte) Builder
	WithFees(fees uint64) Builder
	Now() (Entry, error)
}

// Entry represents an entry
type Entry interface {
	Hash() hash.Hash
	Flag() hash.Hash
	Script() []byte
	Fees() uint64
}
