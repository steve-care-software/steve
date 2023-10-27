package constantvalues

// NewConstantValuesBuilder creates a new constant values builder
func NewConstantValuesBuilder() ConstantValuesBuilder {
	return createConstantValuesBuilder()
}

// NewConstantValueBuilder creates a new constant value builder
func NewConstantValueBuilder() ConstantValueBuilder {
	return createConstantValueBuilder()
}

// ConstantValuesBuilder represents constant values builder
type ConstantValuesBuilder interface {
	Create() ConstantValuesBuilder
	WithList(list []ConstantValue) ConstantValuesBuilder
	Now() (ConstantValues, error)
}

// ConstantValues represents constant values
type ConstantValues interface {
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
	IsVariable() bool
	Variable() string
	IsConstant() bool
	Constant() []byte
}
