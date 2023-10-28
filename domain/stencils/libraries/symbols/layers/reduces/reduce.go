package reduces

import "github.com/steve-care-software/steve/domain/hash"

type reduce struct {
	hash     hash.Hash
	variable string
	length   uint8
}

func createReduce(
	hash hash.Hash,
	variable string,
	length uint8,
) Reduce {
	out := reduce{
		hash:     hash,
		variable: variable,
		length:   length,
	}

	return &out
}

// Hash returns the hash
func (obj *reduce) Hash() hash.Hash {
	return obj.hash
}

// Variable returns the variable
func (obj *reduce) Variable() string {
	return obj.variable
}

// Length returns the length
func (obj *reduce) Length() uint8 {
	return obj.length
}
