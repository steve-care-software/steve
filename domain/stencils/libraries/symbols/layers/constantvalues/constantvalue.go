package constantvalues

import "github.com/steve-care-software/steve/domain/hash"

type constantValue struct {
	hash     hash.Hash
	variable string
	constant []byte
}

func createConstantValueWithVariable(hash hash.Hash, variable string) ConstantValue {
	return createConstantValueInternally(hash, variable, nil)
}

func createConstantValueWithConstant(hash hash.Hash, constant []byte) ConstantValue {
	return createConstantValueInternally(hash, "", constant)
}

func createConstantValueInternally(
	hash hash.Hash,
	variable string,
	constant []byte,
) ConstantValue {
	out := constantValue{
		hash:     hash,
		variable: variable,
		constant: constant,
	}

	return &out
}

// Hash returns the hash
func (obj *constantValue) Hash() hash.Hash {
	return obj.hash
}

// IsVariable returns true if there is a variable, false otherwise
func (obj *constantValue) IsVariable() bool {
	return obj.variable != ""
}

// Variable returns the variable, if any
func (obj *constantValue) Variable() string {
	return obj.variable
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *constantValue) IsConstant() bool {
	return obj.constant != nil
}

// Constant returns the constant, if any
func (obj *constantValue) Constant() []byte {
	return obj.constant
}
