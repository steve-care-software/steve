package layers

type layerInput struct {
	variable string
	layer    Layer
}

func createLayerInput(
	variable string,
	layer Layer,
) LayerInput {
	out := layerInput{
		variable: variable,
		layer:    layer,
	}

	return &out
}

// IsVariable returns true if there is a variable, false otherwise
func (obj *layerInput) IsVariable() bool {
	return obj.variable != ""
}

// Variable returns the variable, if any
func (obj *layerInput) Variable() string {
	return obj.variable
}

// IsLayer returns true if there is a layer, false otherwise
func (obj *layerInput) IsLayer() bool {
	return obj.layer != nil
}

// Layer returns the layer, if any
func (obj *layerInput) Layer() Layer {
	return obj.layer
}
