package instructions

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/assignments"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/operations"
)

type loopCounter struct {
	hash       hash.Hash
	assignment assignments.Assignment
	operation  operations.Operation
	increment  operations.Operation
}

func createLoopCounter(
	hash hash.Hash,
	assignment assignments.Assignment,
	operation operations.Operation,
	increment operations.Operation,
) LoopCounter {
	out := loopCounter{
		hash:       hash,
		assignment: assignment,
		operation:  operation,
		increment:  increment,
	}

	return &out
}

// Hash returns the hash
func (obj *loopCounter) Hash() hash.Hash {
	return obj.hash
}

// Assignment returns the assignment
func (obj *loopCounter) Assignment() assignments.Assignment {
	return obj.assignment
}

// Operation returns the operation
func (obj *loopCounter) Operation() operations.Operation {
	return obj.operation
}

// Increment returns the increment
func (obj *loopCounter) Increment() operations.Operation {
	return obj.increment
}
