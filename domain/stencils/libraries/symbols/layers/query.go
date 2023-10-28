package layers

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/constantvalues"
)

type query struct {
	hash   hash.Hash
	input  constantvalues.ConstantValue
	layer  LayerInput
	values ValueAssignments
}

func createQuery(
	hash hash.Hash,
	input constantvalues.ConstantValue,
	layer LayerInput,
) Query {
	return createQueryInternally(hash, input, layer, nil)
}

func createQueryWithValues(
	hash hash.Hash,
	input constantvalues.ConstantValue,
	layer LayerInput,
	values ValueAssignments,
) Query {
	return createQueryInternally(hash, input, layer, values)
}

func createQueryInternally(
	hash hash.Hash,
	input constantvalues.ConstantValue,
	layer LayerInput,
	values ValueAssignments,
) Query {
	out := query{
		hash:   hash,
		input:  input,
		layer:  layer,
		values: values,
	}

	return &out
}

// Hash returns the hash
func (obj *query) Hash() hash.Hash {
	return obj.hash
}

// Input returns the input
func (obj *query) Input() constantvalues.ConstantValue {
	return obj.input
}

// Layer returns the layer
func (obj *query) Layer() LayerInput {
	return obj.layer
}

// HasValues returns true if there is values, false otherwise
func (obj *query) HasValues() bool {
	return obj.values != nil
}

// Values returns the values, if any
func (obj *query) Values() ValueAssignments {
	return obj.values
}
