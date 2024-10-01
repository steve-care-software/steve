package pointers

import "github.com/steve-care-software/steve/commons/hash"

type pointer struct {
	hash   hash.Hash
	index  uint
	length uint
}

func createPointer(
	hash hash.Hash,
	index uint,
	length uint,
) Pointer {
	out := pointer{
		hash:   hash,
		index:  index,
		length: length,
	}

	return &out
}

// Hash returns the hash
func (obj *pointer) Hash() hash.Hash {
	return obj.hash
}

// Index returns the index
func (obj *pointer) Index() uint {
	return obj.index
}

// Length returns the length
func (obj *pointer) Length() uint {
	return obj.length
}
