package numerics

import "github.com/steve-care-software/steve/commons/hash"

type numeric struct {
	hash hash.Hash
	flag uint8
	size uint8
}

func createNumeric(
	hash hash.Hash,
	flag uint8,
	size uint8,
) Numeric {
	out := numeric{
		hash: hash,
		flag: flag,
		size: size,
	}

	return &out
}

// Hash returns the hash
func (obj *numeric) Hash() hash.Hash {
	return obj.hash
}

// Flag returns the flag
func (obj *numeric) Flag() uint8 {
	return obj.flag
}

// Size returns the size
func (obj *numeric) Size() uint8 {
	return obj.size
}
