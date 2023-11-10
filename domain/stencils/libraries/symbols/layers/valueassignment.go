package layers

import "github.com/steve-care-software/steve/domain/blockchains/hash"

type valueAssignment struct {
	hash  hash.Hash
	name  string
	value Value
}

func createValueAssignment(
	hash hash.Hash,
	name string,
	value Value,
) ValueAssignment {
	out := valueAssignment{
		hash:  hash,
		name:  name,
		value: value,
	}

	return &out
}

// Hash returns the hash
func (obj *valueAssignment) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *valueAssignment) Name() string {
	return obj.name
}

// Value returns the value
func (obj *valueAssignment) Value() Value {
	return obj.value
}
