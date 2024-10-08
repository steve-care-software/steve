package expectations

import "github.com/steve-care-software/steve/hash"

type expectation struct {
	hash   hash.Hash
	path   []string
	isFail bool
}

func createExpectation(
	hash hash.Hash,
	path []string,
	isFail bool,
) Expectation {
	out := expectation{
		hash:   hash,
		path:   path,
		isFail: isFail,
	}

	return &out
}

// Hash returns the hash
func (obj *expectation) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *expectation) Path() []string {
	return obj.path
}

// IsFail returns true if fail, false otherwise
func (obj *expectation) IsFail() bool {
	return obj.isFail
}
