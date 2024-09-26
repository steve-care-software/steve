package instructions

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/instructions/initializations"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/instructions/operations"
)

type loopCounter struct {
	hash           hash.Hash
	initialization initializations.Initialization
	operation      operations.Operation
	increment      operations.Operation
}

func createLoopCounter(
	hash hash.Hash,
	initialization initializations.Initialization,
	operation operations.Operation,
	increment operations.Operation,
) LoopCounter {
	out := loopCounter{
		hash:           hash,
		initialization: initialization,
		operation:      operation,
		increment:      increment,
	}

	return &out
}

// Hash returns the hash
func (obj *loopCounter) Hash() hash.Hash {
	return obj.hash
}

// Initialization returns the initialization
func (obj *loopCounter) Initialization() initializations.Initialization {
	return obj.initialization
}

// Operation returns the operation
func (obj *loopCounter) Operation() operations.Operation {
	return obj.operation
}

// Increment returns the increment
func (obj *loopCounter) Increment() operations.Operation {
	return obj.increment
}
