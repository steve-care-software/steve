package layers

type query struct {
	input  ConstantValue
	layer  LayerInput
	values ValueAssignments
}

func createQuery(
	input ConstantValue,
	layer LayerInput,
) Query {
	return createQueryInternally(input, layer, nil)
}

func createQueryWithValues(
	input ConstantValue,
	layer LayerInput,
	values ValueAssignments,
) Query {
	return createQueryInternally(input, layer, values)
}

func createQueryInternally(
	input ConstantValue,
	layer LayerInput,
	values ValueAssignments,
) Query {
	out := query{
		input:  input,
		layer:  layer,
		values: values,
	}

	return &out
}

// Input returns the input
func (obj *query) Input() ConstantValue {
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
