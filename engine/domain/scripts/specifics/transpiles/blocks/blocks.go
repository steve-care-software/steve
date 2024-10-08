package blocks

import "github.com/steve-care-software/steve/hash"

type blocks struct {
	hash hash.Hash
	list []Block
}

func createBlocks(
	hash hash.Hash,
	list []Block,
) Blocks {
	out := blocks{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *blocks) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *blocks) List() []Block {
	return obj.list
}
