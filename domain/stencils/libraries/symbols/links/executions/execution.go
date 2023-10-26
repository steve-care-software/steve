package executions

import (
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
)

type execution struct {
	layer  layers.LayerInput
	values layers.ValueAssignments
}

func createExecution(
	layer layers.LayerInput,
) Execution {
	return createExecutionInternally(layer, nil)
}

func createExecutionWithValues(
	layer layers.LayerInput,
	values layers.ValueAssignments,
) Execution {
	return createExecutionInternally(layer, values)
}

func createExecutionInternally(
	layer layers.LayerInput,
	values layers.ValueAssignments,
) Execution {
	out := execution{
		layer:  layer,
		values: values,
	}

	return &out
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
