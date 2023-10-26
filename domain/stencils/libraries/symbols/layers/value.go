package layers

type value struct {
	variable string
	constant []byte
	layer    Layer
}

func createValueWithVariable(
	variable string,
) Value {
	return createValueInternally(variable, nil, nil)
}

func createValueWithConstant(
	constant []byte,
) Value {
	return createValueInternally("", constant, nil)
}

func createValueWithLayer(
	layer Layer,
) Value {
	return createValueInternally("", nil, layer)
}

func createValueInternally(
	variable string,
	constant []byte,
	layer Layer,
) Value {
	out := value{
		variable: variable,
		constant: constant,
		layer:    layer,
	}

	return &out
}

// IsVariable returns true if there is a variable, false otherwise
func (obj *value) IsVariable() bool {
	return obj.variable != ""
}

// Variable returns the variable, if any
func (obj *value) Variable() string {
	return obj.variable
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *value) IsConstant() bool {
	return obj.constant != nil
}

// Constant returns the constant, if any
func (obj *value) Constant() []byte {
	return obj.constant
}

// IsLayer returns true if there is a layer, false otherwise
func (obj *value) IsLayer() bool {
	return obj.layer != nil
}

// Layer returns the layer, if any
func (obj *value) Layer() Layer {
	return obj.layer
}
