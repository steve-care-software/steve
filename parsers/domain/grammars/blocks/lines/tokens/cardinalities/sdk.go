package cardinalities

import "github.com/steve-care-software/steve/commons/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a cardinality builder
type Builder interface {
	Create() Builder
	WithMin(min uint) Builder
	WithMax(max uint) Builder
	Now() (Cardinality, error)
}

// Cardinality represents a cardinality
type Cardinality interface {
	Hash() hash.Hash
	Min() uint
	HasMax() bool
	Max() *uint
}
