package parameters

import "github.com/steve-care-software/steve/engine/domain/hash"

type parameters struct {
	hash hash.Hash
	list []Parameter
}

func createParameters(
	hash hash.Hash,
	list []Parameter,
) Parameters {
	out := parameters{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *parameters) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *parameters) List() []Parameter {
	return obj.list
}
