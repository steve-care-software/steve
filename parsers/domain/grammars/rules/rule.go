package rules

import "github.com/steve-care-software/steve/engine/domain/hash"

type rule struct {
	hash  hash.Hash
	name  string
	bytes []byte
}

func createRule(
	hash hash.Hash,
	name string,
	bytes []byte,
) Rule {
	out := rule{
		hash:  hash,
		name:  name,
		bytes: bytes,
	}

	return &out
}

// Hash returns the hash
func (obj *rule) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *rule) Name() string {
	return obj.name
}

// Bytes returns the bytes
func (obj *rule) Bytes() []byte {
	return obj.bytes
}
