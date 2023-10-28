package constantvalues

import "github.com/steve-care-software/steve/domain/hash"

// NewConstantValuesBuilder creates a new constant values builder
func NewConstantValuesBuilder() ConstantValuesBuilder {
	hashAdapter := hash.NewAdapter()
	return createConstantValuesBuilder(
		hashAdapter,
	)
}

// NewConstantValueBuilder creates a new constant value builder
func NewConstantValueBuilder() ConstantValueBuilder {
	hashAdapter := hash.NewAdapter()
	return createConstantValueBuilder(
		hashAdapter,
	)
}

// ConstantValuesBuilder represents constant values builder
type ConstantValuesBuilder interface {
	Create() ConstantValuesBuilder
	WithList(list []ConstantValue) ConstantValuesBuilder
	Now() (ConstantValues, error)
}

// ConstantValues represents constant values
type ConstantValues interface {
	Hash() hash.Hash
	List() []ConstantValue
}

// ConstantValueBuilder represents a constant value builder
type ConstantValueBuilder interface {
	Create() ConstantValueBuilder
	WithVariable(variable string) ConstantValueBuilder
	WithConstant(constant []byte) ConstantValueBuilder
	Now() (ConstantValue, error)
}

// ConstantValue represents a constant value
type ConstantValue interface {
	Hash() hash.Hash
	IsVariable() bool
	Variable() string
	IsConstant() bool
	Constant() []byte
}

// Service represents a symbol service
type Service interface {
	Insert(context uint, container []string, value constantValue) error
	InsertList(context uint, container []string, values constantValues) error
}
