package constantvalues

import "github.com/steve-care-software/steve/domain/blockchains/hash"

type constantValues struct {
	hash hash.Hash
	list []ConstantValue
}

func createConstantValues(
	hash hash.Hash,
	list []ConstantValue,
) ConstantValues {
	out := constantValues{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *constantValues) Hash() hash.Hash {
	return obj.hash
}

// List returns the constantValues
func (obj *constantValues) List() []ConstantValue {
	return obj.list
}
