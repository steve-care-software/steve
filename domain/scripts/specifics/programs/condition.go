package programs

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/operations"
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
