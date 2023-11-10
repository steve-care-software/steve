package parameters

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/parameters/kinds"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewParameterBuilder creates a new parameter builder
func NewParameterBuilder() ParameterBuilder {
	hashAdapter := hash.NewAdapter()
	return createParameterBuilder(
		hashAdapter,
	)
}

// Builder represents a parameters builder
type Builder interface {
	Create() Builder
	WithList(list []Parameter) Builder
	Now() (Parameters, error)
}

// Parameters represents a parameters
type Parameters interface {
	Hash() hash.Hash
	List() []Parameter
}

// ParameterBuilder represents a parameter builder
type ParameterBuilder interface {
	Create() ParameterBuilder
	WithName(name string) ParameterBuilder
	WithKind(kind kinds.Kind) ParameterBuilder
	Now() (Parameter, error)
}

// Parameter represents a parameter
type Parameter interface {
	Hash() hash.Hash
	Name() string
	Kind() kinds.Kind
}
