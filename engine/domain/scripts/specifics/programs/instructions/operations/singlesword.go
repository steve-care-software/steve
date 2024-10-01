package operations

import "github.com/steve-care-software/steve/commons/hash"

type singleSword struct {
	hash     hash.Hash
	variable string
	flag     uint8
}

func createSingleSword(
	hash hash.Hash,
	variable string,
	flag uint8,
) SingleSword {
	out := singleSword{
		hash:     hash,
		variable: variable,
		flag:     flag,
	}

	return &out
}

// Hash returns the hash
func (obj *singleSword) Hash() hash.Hash {
	return obj.hash
}

// Variable returns the variable
func (obj *singleSword) Variable() string {
	return obj.variable
}

// Flag returns the flag
func (obj *singleSword) Flag() uint8 {
	return obj.flag
}
