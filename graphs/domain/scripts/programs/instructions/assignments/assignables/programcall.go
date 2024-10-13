package assignables

import "github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references"

type programCall struct {
	references references.References
	params     MapKeyValues
}

func createProgramCall(references references.References) ProgramCall {
	return createProgramCallInternally(references, nil)
}

func createProgramCallWithParams(references references.References, params MapKeyValues) ProgramCall {
	return createProgramCallInternally(references, params)
}

func createProgramCallInternally(references references.References, params MapKeyValues) ProgramCall {
	return &programCall{
		references: references,
		params:     params,
	}
}

// References returns the references of the program call
func (obj *programCall) References() references.References {
	return obj.references
}

// HasParams returns true if the program call has parameters
func (obj *programCall) HasParams() bool {
	return obj.params != nil
}

// Params returns the parameters of the program call if present
func (obj *programCall) Params() MapKeyValues {
	return obj.params
}
