package assignables

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/instructions/calls"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/instructions/operations"
)

type assignable struct {
	hash      hash.Hash
	operation operations.Operation
	call      calls.Call
}

func createAssignableWithOperation(
	hash hash.Hash,
	operation operations.Operation,
) Assignable {
	return createAssignableInternally(hash, operation, nil)
}

func createAssignableWithCall(
	hash hash.Hash,
	call calls.Call,
) Assignable {
	return createAssignableInternally(hash, nil, call)
}

func createAssignableInternally(
	hash hash.Hash,
	operation operations.Operation,
	call calls.Call,
) Assignable {
	out := assignable{
		hash:      hash,
		operation: operation,
		call:      call,
	}

	return &out
}

// Hash returns the hash
func (obj *assignable) Hash() hash.Hash {
	return obj.hash
}

// IsOperation returns true if there is an operation, false otherwise
func (obj *assignable) IsOperation() bool {
	return obj.operation != nil
}

// Operation returns the operation, if any
func (obj *assignable) Operation() operations.Operation {
	return obj.operation
}

// IsCall returns true if there is a call, false otherwise
func (obj *assignable) IsCall() bool {
	return obj.call != nil
}

// Call returns the call, if any
func (obj *assignable) Call() calls.Call {
	return obj.call
}
