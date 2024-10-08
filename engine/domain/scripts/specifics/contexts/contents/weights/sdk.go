package weights

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/components/suites"
	"github.com/steve-care-software/steve/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewWeightBuilder creates a weight builder
func NewWeightBuilder() WeightBuilder {
	hashAdapter := hash.NewAdapter()
	return createWeightBuilder(
		hashAdapter,
	)
}

// Builder represents the weights builder
type Builder interface {
	Create() Builder
	WithList(list []Weight) Builder
	Now() (Weights, error)
}

// Weights represents weights
type Weights interface {
	Hash() hash.Hash
	List() []Weight
}

// WeightBuilder represents a weight builder
type WeightBuilder interface {
	Create() WeightBuilder
	WithName(name string) WeightBuilder
	WithValue(value uint) WeightBuilder
	WithReverse(reverse string) WeightBuilder
	WithSuites(suites suites.Suites) WeightBuilder
	Now() (Weight, error)
}

// Weight represents a weight
type Weight interface {
	Hash() hash.Hash
	Name() string
	Value() uint
	HasReverse() bool
	Reverse() string
	HasSuites() bool
	Suites() suites.Suites
}
