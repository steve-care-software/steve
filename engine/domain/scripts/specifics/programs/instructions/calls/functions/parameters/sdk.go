package parameters

import "github.com/steve-care-software/steve/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewParameterBuilder creates a new parameter builder
func NewParameterBuilder() ParameterBuilder {
	hashAdapter := hash.NewAdapter()
	return createPArameterBuilder(
		hashAdapter,
	)
}

// Builder represents the parameters builder
type Builder interface {
	Create() Builder
	WithList(list []Parameter) Builder
	Now() (Parameters, error)
}

// Parameters represents a func call parameters
type Parameters interface {
	Hash() hash.Hash
	List() []Parameter
}

// ParameterBuilder represents the parameter builder
type ParameterBuilder interface {
	Create() ParameterBuilder
	WithCurrent(current string) ParameterBuilder
	WithLocal(local string) ParameterBuilder
	Now() (Parameter, error)
}

// Parameter represents a func call parameter
type Parameter interface {
	Hash() hash.Hash
	Current() string
	Local() string
}
