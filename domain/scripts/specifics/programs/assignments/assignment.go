package assignments

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/operations"
)

type assignment struct {
	hash      hash.Hash
	variables []string
	operation operations.Operation
	isInitial bool
}

func createAssignment(
	hash hash.Hash,
	variables []string,
	operation operations.Operation,
	isInitial bool,
) Assignment {
	out := assignment{
		hash:      hash,
		variables: variables,
		operation: operation,
		isInitial: isInitial,
	}

	return &out
}

// Hash returns the hash
func (obj *assignment) Hash() hash.Hash {
	return obj.hash
}

// Variables returns the variables
func (obj *assignment) Variables() []string {
	return obj.variables
}

// Operation returns the operation
func (obj *assignment) Operation() operations.Operation {
	return obj.operation
}

// IsInitial returns true if initial, false otherwise
func (obj *assignment) IsInitial() bool {
	return obj.isInitial
}
