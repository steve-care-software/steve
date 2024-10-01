package suites

import "github.com/steve-care-software/steve/commons/hash"

type suite struct {
	hash   hash.Hash
	name   string
	input  []byte
	output []byte
}

func createSuite(
	hash hash.Hash,
	name string,
	input []byte,
	output []byte,
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
func (obj *suite) Output() []byte {
	return obj.output
}
