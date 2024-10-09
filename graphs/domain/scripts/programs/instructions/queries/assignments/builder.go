package assignments

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references/externals"
)

type builder struct {
	external externals.External
	variable string
}

func createBuilder() Builder {
	return &builder{
		external: nil,
		variable: "",
	}
}

// Create initializes the assignment builder
func (obj *builder) Create() Builder {
	return createBuilder()
}

// WithExternal adds an external to the assignment builder
func (obj *builder) WithExternal(external externals.External) Builder {
	obj.external = external
	return obj
}

// WithVariable adds a variable to the assignment builder
func (obj *builder) WithVariable(variable string) Builder {
	obj.variable = variable
	return obj
}

// Now builds a new Assignment instance
func (obj *builder) Now() (Assignment, error) {
	if obj.external == nil {
		return nil, errors.New("the external is mandatory in order to build an Assignment instance")
	}

	if obj.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build an Assignment instance")
	}

	return createAssignment(obj.external, obj.variable), nil
}
