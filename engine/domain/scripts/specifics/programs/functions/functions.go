package functions

import "github.com/steve-care-software/steve/engine/domain/hash"

type functions struct {
	hash hash.Hash
	list []Function
}

func createFunctions(
	hash hash.Hash,
	list []Function,
) Functions {
	out := functions{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *functions) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *functions) List() []Function {
	return obj.list
}
