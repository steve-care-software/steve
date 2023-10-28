package pointers

import "github.com/steve-care-software/steve/domain/hash"

type pointers struct {
	hash hash.Hash
	list []Pointer
}

func createPointers(
	hash hash.Hash,
	list []Pointer,
) Pointers {
	out := pointers{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *pointers) Hash() hash.Hash {
	return obj.hash
}

// List returns the pointers
func (obj *pointers) List() []Pointer {
	return obj.list
}
