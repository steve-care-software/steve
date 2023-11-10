package suites

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns"
)

type suite struct {
	hash   hash.Hash
	name   string
	input  []byte
	output returns.Return
}

func createSuite(
	hash hash.Hash,
	name string,
	input []byte,
	output returns.Return,
) Suite {
	out := suite{
		hash:   hash,
		name:   name,
		input:  input,
		output: output,
	}

	return &out
}

// Hash returns the hash
func (obj *suite) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *suite) Name() string {
	return obj.name
}

// Input returns the input
func (obj *suite) Input() []byte {
	return obj.input
}

// Output returns the output
func (obj *suite) Output() returns.Return {
	return obj.output
}
