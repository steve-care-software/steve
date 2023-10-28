package executions

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
)

type execution struct {
	hash   hash.Hash
	layer  layers.LayerInput
	values layers.ValueAssignments
}

func createExecution(
	hash hash.Hash,
	layer layers.LayerInput,
) Execution {
	return createExecutionInternally(hash, layer, nil)
}

func createExecutionWithValues(
	hash hash.Hash,
	layer layers.LayerInput,
	values layers.ValueAssignments,
) Execution {
	return createExecutionInternally(hash, layer, values)
}

func createExecutionInternally(
	hash hash.Hash,
	layer layers.LayerInput,
	values layers.ValueAssignments,
) Execution {
	out := execution{
		hash:   hash,
		layer:  layer,
		values: values,
	}

	return &out
}

// Hash returns the hash
func (obj *execution) Hash() hash.Hash {
	return obj.hash
}

// Layer returns the layer
func (obj *execution) Layer() layers.LayerInput {
	return obj.layer
}

// HasValues returns true if there is values, false otherwise
func (obj *execution) HasValues() bool {
	return obj.values != nil
}

// Values returns values, if any
func (obj *execution) Values() layers.ValueAssignments {
	return obj.values
}
