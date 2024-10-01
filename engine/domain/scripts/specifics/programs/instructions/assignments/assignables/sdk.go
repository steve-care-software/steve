package assignables

import (
	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/calls"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/operations"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents an assignable builder
type Builder interface {
	Create() Builder
	WithOperation(operation operations.Operation) Builder
	WithCall(call calls.Call) Builder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	Hash() hash.Hash
	IsOperation() bool
	Operation() operations.Operation
	IsCall() bool
	Call() calls.Call
}
