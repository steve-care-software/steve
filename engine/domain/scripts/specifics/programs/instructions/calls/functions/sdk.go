package functions

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/calls/functions/parameters"
	"github.com/steve-care-software/steve/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the function builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithParameters(parameters parameters.Parameters) Builder
	IsEngine() Builder
	Now() (Function, error)
}

// Function represents a func call
type Function interface {
	Hash() hash.Hash
	Name() string
	Parameters() parameters.Parameters
	IsEngine() bool
}
