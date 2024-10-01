package suites

import (
	"github.com/steve-care-software/steve/engine/domain/hash"
)

type suite struct {
	hash   hash.Hash
	name   string
	input  []byte
	isFail bool
}

func createSuite(
	hash hash.Hash,
	name string,
	input []byte,
	isFail bool,
) Suite {
	out := suite{
		hash:   hash,
		name:   name,
		input:  input,
		isFail: isFail,
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

// IsFail returns true if expected to fail, false otherwise
func (obj *suite) IsFail() bool {
	return obj.isFail
}
