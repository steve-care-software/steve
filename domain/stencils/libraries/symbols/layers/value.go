package layers

import "github.com/steve-care-software/steve/domain/blockchains/hash"

type value struct {
	hash     hash.Hash
	variable string
	constant []byte
	layer    Layer
}

func createValueWithVariable(
	hash hash.Hash,
	variable string,
) Value {
	return createValueInternally(hash, variable, nil, nil)
}

func createValueWithConstant(
	hash hash.Hash,
	constant []byte,
) Value {
	return createValueInternally(hash, "", constant, nil)
}

func createValueWithLayer(
	hash hash.Hash,
	layer Layer,
) Value {
	return createValueInternally(hash, "", nil, layer)
}

func createValueInternally(
	hash hash.Hash,
	variable string,
	constant []byte,
	layer Layer,
) Value {
	out := value{
		hash:     hash,
		variable: variable,
		constant: constant,
		layer:    layer,
	}

	return &out
}

// Hash returns the hash
func (obj *value) Hash() hash.Hash {
	return obj.hash
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
