package assignments

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/containers"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/assignments/assignables"
	"github.com/steve-care-software/steve/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the assignment builder
type Builder interface {
	Create() Builder
	WithVariables(variables []string) Builder
	WithAssignable(assignable assignables.Assignable) Builder
	WithInitial(initial containers.Container) Builder
	Now() (Assignment, error)
}

// Assignment represents an assignment
type Assignment interface {
	Hash() hash.Hash
	Variables() []string
	Assignable() assignables.Assignable
	HasInitial() bool
	Initial() containers.Container
}
