package operations

import "github.com/steve-care-software/steve/commons/hash"

type standard struct {
	hash   hash.Hash
	first  Operation
	second Operation
	flag   uint8
}

func createStandard(
	hash hash.Hash,
	first Operation,
	second Operation,
	flag uint8,
) Standard {
	out := standard{
		hash:   hash,
		first:  first,
		second: second,
		flag:   flag,
	}

	return &out
}

// Hash returns the hash
func (obj *standard) Hash() hash.Hash {
	return obj.hash
}

// First returns the first operation
func (obj *standard) First() Operation {
	return obj.first
}

// Second returns the second operation
func (obj *standard) Second() Operation {
	return obj.second
}

// Flag returns the flag
func (obj *standard) Flag() uint8 {
	return obj.flag
}
