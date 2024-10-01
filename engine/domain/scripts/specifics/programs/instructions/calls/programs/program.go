package programs

import "github.com/steve-care-software/steve/commons/hash"

type program struct {
	hash  hash.Hash
	name  string
	input string
}

func createProgram(
	hash hash.Hash,
	name string,
	input string,
) Program {
	out := program{
		hash:  hash,
		name:  name,
		input: input,
	}

	return &out
}

// Hash returns the hash
func (obj *program) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *program) Name() string {
	return obj.name
}

// Input returns the input
func (obj *program) Input() string {
	return obj.input
}
