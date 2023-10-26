package layers

type constantValue struct {
	variable string
	constant []byte
}

func createConstantValueWithVariable(variable string) ConstantValue {
	return createConstantValueInternally(variable, nil)
}

func createConstantValueWithConstant(constant []byte) ConstantValue {
	return createConstantValueInternally("", constant)
}

func createConstantValueInternally(
	variable string,
	constant []byte,
) ConstantValue {
	out := constantValue{
		variable: variable,
		constant: constant,
	}

	return &out
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
