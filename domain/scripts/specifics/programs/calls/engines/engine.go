package engines

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/calls/functions"
)

type engine struct {
	hash     hash.Hash
	scope    uint8
	function functions.Function
}

func createEngine(
	hash hash.Hash,
	scope uint8,
	function functions.Function,
) Engine {
	out := engine{
		hash:     hash,
		scope:    scope,
		function: function,
	}

	return &out
}

// Hash returns the hash
func (obj *engine) Hash() hash.Hash {
	return obj.hash
}

// Scope returns the scope
func (obj *engine) Scope() uint8 {
	return obj.scope
}

// Function returns the function
func (obj *engine) Function() functions.Function {
	return obj.function
}
