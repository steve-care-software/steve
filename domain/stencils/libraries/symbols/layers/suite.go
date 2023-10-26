package layers

import (
	result_returns "github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns"
)

type suite struct {
	name   string
	input  []byte
	ret    result_returns.Return
	values ValueAssignments
}

func createSuite(
	name string,
	input []byte,
	ret result_returns.Return,
) Suite {
	return createSuiteInternally(name, input, ret, nil)
}

func createSuiteWithValues(
	name string,
	input []byte,
	ret result_returns.Return,
	values ValueAssignments,
) Suite {
	return createSuiteInternally(name, input, ret, values)
}

func createSuiteInternally(
	name string,
	input []byte,
	ret result_returns.Return,
	values ValueAssignments,
) Suite {
	out := suite{
		name:   name,
		input:  input,
		ret:    ret,
		values: values,
	}

	return &out
}

// Name returns the name
func (obj *suite) Name() string {
	return obj.name
}

// Input returns the input
func (obj *suite) Input() []byte {
	return obj.input
}

// Return returns the return value
func (obj *suite) Return() result_returns.Return {
	return obj.ret
}

// HasValues returns true if there is values, false otherwise
func (obj *suite) HasValues() bool {
	return obj.values != nil
}

// Values returns the values, if any
func (obj *suite) Values() ValueAssignments {
	return obj.values
}
