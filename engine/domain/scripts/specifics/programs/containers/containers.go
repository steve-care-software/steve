package containers

import "github.com/steve-care-software/steve/commons/hash"

type containers struct {
	hash hash.Hash
	list []Container
}

func createContainers(
	hash hash.Hash,
	list []Container,
) Containers {
	out := containers{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *containers) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *containers) List() []Container {
	return obj.list
}
