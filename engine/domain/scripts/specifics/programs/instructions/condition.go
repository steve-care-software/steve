package instructions

import (
	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/operations"
)

type condition struct {
	hash         hash.Hash
	operation    operations.Operation
	instructions Instructions
}

func createCondition(
	hash hash.Hash,
	operation operations.Operation,
	instructions Instructions,
) Condition {
	out := condition{
		hash:         hash,
		operation:    operation,
		instructions: instructions,
	}

	return &out
}

// Hash returns the hash
func (obj *condition) Hash() hash.Hash {
	return obj.hash
}

// Operation returns the operation
func (obj *condition) Operation() operations.Operation {
	return obj.operation
}

// Instructions returns the instructions
func (obj *condition) Instructions() Instructions {
	return obj.instructions
}
