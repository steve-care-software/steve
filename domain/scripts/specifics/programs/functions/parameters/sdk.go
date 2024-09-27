package parameters

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/containers"
)

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

// Parameters represents func parameters
type Parameters interface {
	Hash() hash.Hash
	List() []Parameter
}

// ParameterBuilder represents a parameter builder
type ParameterBuilder interface {
	Create() ParameterBuilder
	WithName(name string) ParameterBuilder
	WithContainer(container containers.Container) ParameterBuilder
	IsMandatory() ParameterBuilder
	Now() (Parameter, error)
}

// Parameter represents a func parameter
type Parameter interface {
	Hash() hash.Hash
	Name() string
	Container() containers.Container
	IsMandatory() bool
}
