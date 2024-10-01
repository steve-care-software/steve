package operations

import (
	"github.com/steve-care-software/steve/engine/domain/hash"
)

type operation struct {
	hash        hash.Hash
	standard    Standard
	singleSword SingleSword
	bitshift    BitShift
	variable    string
	value       any
}

func createOperationWithStandard(
	hash hash.Hash,
	standard Standard,
) Operation {
	return createOperationInternally(hash, standard, nil, nil, "", nil)
}

func createOperationWithSingleSword(
	hash hash.Hash,
	singleSword SingleSword,
) Operation {
	return createOperationInternally(hash, nil, singleSword, nil, "", nil)
}

func createOperationWithBitShift(
	hash hash.Hash,
	bitshift BitShift,
) Operation {
	return createOperationInternally(hash, nil, nil, bitshift, "", nil)
}

func createOperationWithVariable(
	hash hash.Hash,
	variable string,
) Operation {
	return createOperationInternally(hash, nil, nil, nil, variable, nil)
}

func createOperationWithValue(
	hash hash.Hash,
	value any,
) Operation {
	return createOperationInternally(hash, nil, nil, nil, "", value)
}

func createOperationInternally(
	hash hash.Hash,
	standard Standard,
	singleSword SingleSword,
	bitshift BitShift,
	variable string,
	value any,
) Operation {
	out := operation{
		hash:        hash,
		standard:    standard,
		singleSword: singleSword,
		bitshift:    bitshift,
		variable:    variable,
		value:       value,
	}

	return &out
}

// Hash returns the hash
func (obj *operation) Hash() hash.Hash {
	return obj.hash
}

// IsStandard returns true if there is the standard, false otherwise
func (obj *operation) IsStandard() bool {
	return obj.standard != nil
}

// Standard returns the standard, if any
func (obj *operation) Standard() Standard {
	return obj.standard
}

// IsSingleSword returns true if there is the singleSword, false otherwise
func (obj *operation) IsSingleSword() bool {
	return obj.singleSword != nil
}

// SingleSword returns the singleSword, if any
func (obj *operation) SingleSword() SingleSword {
	return obj.singleSword
}

// IsBitShift returns true if there is a bitshift, false otherwise
func (obj *operation) IsBitShift() bool {
	return obj.bitshift != nil
}

// BitShift returns the bitshift, if any
func (obj *operation) BitShift() BitShift {
	return obj.bitshift
}

// IsVariable returns true if there is a variable, false otherwise
func (obj *operation) IsVariable() bool {
	return obj.variable != ""
}

// Variable returns the variable, if any
func (obj *operation) Variable() string {
	return obj.variable
}

// IsValue returns true if there is a value, false otherwise
func (obj *operation) IsValue() bool {
	return obj.value != nil
}

// Value returns the value, if any
func (obj *operation) Value() any {
	return obj.value
}
