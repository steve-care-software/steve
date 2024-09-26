package assignments

import (
	"github.com/steve-care-software/steve/domain/hash"
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
	WithOperation(operation operations.Operation) Builder
	IsInitial() Builder
	Now() (Assignment, error)
}

// Assignment represents an assignment
type Assignment interface {
	Hash() hash.Hash
	Variables() []string
	Operation() operations.Operation
	IsInitial() bool
}
