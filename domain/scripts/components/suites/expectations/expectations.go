package expectations

import "github.com/steve-care-software/steve/domain/hash"

type expectations struct {
	hash hash.Hash
	list []Expectation
}

func createExpectations(
	hash hash.Hash,
	list []Expectation,
) Expectations {
	out := expectations{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *expectations) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *expectations) List() []Expectation {
	return obj.list
}
