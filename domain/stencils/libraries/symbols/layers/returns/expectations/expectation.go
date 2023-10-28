package expectations

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns/kinds"
)

type expectation struct {
	hash     hash.Hash
	variable string
	kind     kinds.Kind
}

func createExpectation(
	hash hash.Hash,
	variable string,
	kind kinds.Kind,
) Expectation {
	out := expectation{
		hash:     hash,
		variable: variable,
		kind:     kind,
	}

	return &out
}

// Hash returns the hash
func (obj *expectation) Hash() hash.Hash {
	return obj.hash
}

// Variable returns the variable
func (obj *expectation) Variable() string {
	return obj.variable
}

// Kind returns the kind
func (obj *expectation) Kind() kinds.Kind {
	return obj.kind
}
