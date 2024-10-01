package contexts

import "github.com/steve-care-software/steve/engine/domain/hash"

type contexts struct {
	hash hash.Hash
	list []Context
}

func createContexts(
	hash hash.Hash,
	list []Context,
) Contexts {
	out := contexts{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *contexts) Hash() hash.Hash {
	return obj.hash
}

// List returns the list of contexts
func (obj *contexts) List() []Context {
	return obj.list
}
