package weights

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/contexts/contents/suites"
)

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
