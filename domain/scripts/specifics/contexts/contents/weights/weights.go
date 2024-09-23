package weights

import "github.com/steve-care-software/steve/domain/hash"

type weights struct {
	hash hash.Hash
	list []Weight
}

func createWeights(
	hash hash.Hash,
	list []Weight,
) Weights {
	out := weights{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *weights) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *weights) List() []Weight {
	return obj.list
}
