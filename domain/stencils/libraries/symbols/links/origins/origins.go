package origins

import "github.com/steve-care-software/steve/domain/blockchains/hash"

type origins struct {
	hash hash.Hash
	list []Origin
}

func createOrigins(
	hash hash.Hash,
	list []Origin,
) Origins {
	out := origins{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *origins) Hash() hash.Hash {
	return obj.hash
}

// List returns the origins
func (obj *origins) List() []Origin {
	return obj.list
}
