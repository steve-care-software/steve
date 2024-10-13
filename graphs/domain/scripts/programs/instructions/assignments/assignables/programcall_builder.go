package assignables

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references"
)

type programCallBuilder struct {
	references references.References
	params     MapKeyValues
}

func createProgramCallBuilder() ProgramCallBuilder {
	return &programCallBuilder{
		references: nil,
		params:     nil,
	}
}

// Create initializes the program call builder
func (obj *programCallBuilder) Create() ProgramCallBuilder {
	return createProgramCallBuilder()
}

// WithReferences adds references to the builder
func (obj *programCallBuilder) WithReferences(references references.References) ProgramCallBuilder {
	obj.references = references
	return obj
}

// WithParams adds parameters to the builder
func (obj *programCallBuilder) WithParams(params MapKeyValues) ProgramCallBuilder {
	obj.params = params
	return obj
}

// Now builds a new ProgramCall instance
func (obj *programCallBuilder) Now() (ProgramCall, error) {
	if obj.references == nil {
		return nil, errors.New("the references are mandatory in order to build a ProgramCall instance")
	}

	if obj.params != nil {
		return createProgramCallWithParams(obj.references, obj.params), nil
	}

	return createProgramCall(obj.references), nil
}
