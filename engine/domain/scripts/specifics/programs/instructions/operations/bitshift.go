package operations

import "github.com/steve-care-software/steve/commons/hash"

type bitshift struct {
	hash      hash.Hash
	operation Operation
	flag      uint8
	amount    uint8
}

func createBitshift(
	hash hash.Hash,
	operation Operation,
	flag uint8,
	amount uint8,
) BitShift {
	out := bitshift{
		hash:      hash,
		operation: operation,
		flag:      flag,
		amount:    amount,
	}

	return &out
}

// Hash returns the hash
func (obj *bitshift) Hash() hash.Hash {
	return obj.hash
}

// Operation returns the operation
func (obj *bitshift) Operation() Operation {
	return obj.operation
}

// Flag returns the flag
func (obj *bitshift) Flag() uint8 {
	return obj.flag
}

// Amount returns the amount
func (obj *bitshift) Amount() uint8 {
	return obj.amount
}
