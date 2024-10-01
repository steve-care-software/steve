package functions

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/calls/functions/parameters"
)

type function struct {
	hash       hash.Hash
	name       string
	parameters parameters.Parameters
	isEngine   bool
}

func createFunction(
	hash hash.Hash,
	name string,
	parameters parameters.Parameters,
	isEngine bool,
) Function {
	out := function{
		hash:       hash,
		name:       name,
		parameters: parameters,
		isEngine:   isEngine,
	}

	return &out
}

// Hash returns the hash
func (obj *function) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *function) Name() string {
	return obj.name
}

// Parameters returns the parameters
func (obj *function) Parameters() parameters.Parameters {
	return obj.parameters
}

// IsEngine returns true if there is an engine, false otherwise
func (obj *function) IsEngine() bool {
	return obj.isEngine
}
