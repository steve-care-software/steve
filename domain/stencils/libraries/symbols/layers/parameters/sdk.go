package parameters

import "github.com/steve-care-software/steve/domain/pointers/symbols"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Validate validates the kind
func Validate(kind uint8) bool {
	return kind&symbols.KindBytes|symbols.KindLayer == 0
}

// NewParameterBuilder creates a new parameter builder
func NewParameterBuilder() ParameterBuilder {
	return createParameterBuilder()
}

// Builder represents a parameters builder
type Builder interface {
	Create() Builder
	WithList(list []Parameter) Builder
	Now() (Parameters, error)
}

// Parameters represents a parameters
type Parameters interface {
	List() []Parameter
}

// ParameterBuilder represents a parameter builder
type ParameterBuilder interface {
	Create() ParameterBuilder
	WithName(name string) ParameterBuilder
	WithKind(kind uint8) ParameterBuilder
	Now() (Parameter, error)
}

// Parameter represents a parameter
type Parameter interface {
	Name() string
	Kind() uint8
}
