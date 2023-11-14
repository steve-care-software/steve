package reduces

import (
	"github.com/steve-care-software/steve/vms/bytes/results/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a reduce
type Builder interface {
	Create() Builder
	WithVariable(variable string) Builder
	WithLength(length uint8) Builder
	Now() (Reduce, error)
}

// Reduce represents a reduce
type Reduce interface {
	Hash() hash.Hash
	Variable() string
	Length() uint8
}
