package assignments

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/containers"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/instructions/calls"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/instructions/operations"
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
	WithAssignable(assignable Assignable) Builder
	WithInitial(initial containers.Container) Builder
	Now() (Assignment, error)
}

// Assignment represents an assignment
type Assignment interface {
	Hash() hash.Hash
	Variables() []string
	Assignable() Assignable
	HasInitial() bool
	Initial() containers.Container
}

// Assignable represents an assignable
type Assignable interface {
	Hash() hash.Hash
	IsOperation() bool
	Operation() operations.Operation
	IsCall() bool
	Call() calls.Call
}
