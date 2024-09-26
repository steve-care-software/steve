package initializations

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/containers"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/instructions/assignments"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithContainer(container containers.Container) Builder
	WithAssignment(assignment assignments.Assignment) Builder
	Now() (Initialization, error)
}

// Initialization represents a variable initialization
type Initialization interface {
	Hash() hash.Hash
	Container() containers.Container
	Assignment() assignments.Assignment
}
