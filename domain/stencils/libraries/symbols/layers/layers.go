package layers

import "github.com/steve-care-software/steve/domain/blockchains/hash"

type layers struct {
	hash hash.Hash
	list []Layer
}

func createLayers(
	hash hash.Hash,
	list []Layer,
) Layers {
	out := layers{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *layers) Hash() hash.Hash {
	return obj.hash
}

// List returns the layers
func (obj *layers) List() []Layer {
	return obj.list
}
