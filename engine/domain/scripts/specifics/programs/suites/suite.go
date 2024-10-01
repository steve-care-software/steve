package suites

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions"
)

type suite struct {
	hash        hash.Hash
	init        instructions.Instructions
	input       []byte
	expectation []byte
}

func createSuite(
	hash hash.Hash,
	init instructions.Instructions,
	input []byte,
	expectation []byte,
) Suite {
	out := suite{
		hash:        hash,
		init:        init,
		input:       input,
		expectation: expectation,
	}

	return &out
}

// Hash returns the hash
func (obj *suite) Hash() hash.Hash {
	return obj.hash
}

// Init returns the init
func (obj *suite) Init() instructions.Instructions {
	return obj.init
}

// Input returns the input
func (obj *suite) Input() []byte {
	return obj.input
}

// Expectation returns the expectation
func (obj *suite) Expectation() []byte {
	return obj.expectation
}
