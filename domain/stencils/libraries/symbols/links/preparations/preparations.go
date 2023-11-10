package preparations

import "github.com/steve-care-software/steve/domain/blockchains/hash"

type preparations struct {
	hash hash.Hash
	list []Preparation
}

func createPreparations(
	hash hash.Hash,
	list []Preparation,
) Preparations {
	out := preparations{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *preparations) Hash() hash.Hash {
	return obj.hash
}

// List returns the preparations
func (obj *preparations) List() []Preparation {
	return obj.list
}
