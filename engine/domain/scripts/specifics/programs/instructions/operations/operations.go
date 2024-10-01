package operations

import "github.com/steve-care-software/steve/engine/domain/hash"

type operations struct {
	hash hash.Hash
	list []Operation
}

func createOperations(
	hash hash.Hash,
	list []Operation,
) Operations {
	out := operations{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *operations) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *operations) List() []Operation {
	return obj.list
}
