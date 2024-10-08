package instructions

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/operations"
	"github.com/steve-care-software/steve/hash"
)

type loopKeyValue struct {
	hash      hash.Hash
	keyName   string
	valueName string
	operation operations.Operation
}

func createLoopKeyValue(
	hash hash.Hash,
	keyName string,
	valueName string,
	operation operations.Operation,
) LoopKeyValue {
	out := loopKeyValue{
		hash:      hash,
		keyName:   keyName,
		valueName: valueName,
		operation: operation,
	}

	return &out
}

// Hash returns the hash
func (obj *loopKeyValue) Hash() hash.Hash {
	return obj.hash
}

// KeyName returns the keyName
func (obj *loopKeyValue) KeyName() string {
	return obj.keyName
}

// ValueName returns the valueName
func (obj *loopKeyValue) ValueName() string {
	return obj.valueName
}

// Operation returns the operation
func (obj *loopKeyValue) Operation() operations.Operation {
	return obj.operation
}
