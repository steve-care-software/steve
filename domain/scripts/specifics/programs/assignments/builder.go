package assignments

import (
	"errors"
	"strings"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/operations"
)

type builder struct {
	hashAapter hash.Adapter
	variables  []string
	operation  operations.Operation
	isInitial  bool
}

func createBuilder(
	hashAapter hash.Adapter,
) Builder {
	out := builder{
		hashAapter: hashAapter,
		variables:  nil,
		operation:  nil,
		isInitial:  false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAapter,
	)
}

// WithVariables add variables to the builder
func (app *builder) WithVariables(variables []string) Builder {
	app.variables = variables
	return app
}

// WithOperation add opration to the builder
func (app *builder) WithOperation(operation operations.Operation) Builder {
	app.operation = operation
	return app
}

// IsInitial flags the builder as initial
func (app *builder) IsInitial() Builder {
	app.isInitial = true
	return app
}

// Now builds a  new Assignment instance
func (app *builder) Now() (Assignment, error) {
	if app.variables != nil && len(app.variables) <= 0 {
		app.variables = nil
	}

	if app.variables == nil {
		return nil, errors.New("the variables is mandatory in order to build an Assignment instance")
	}

	if app.operation == nil {
		return nil, errors.New("the operation is mandatory in order to build an Assignment instance")
	}

	isInitial := "false"
	if app.isInitial {
		isInitial = "true"
	}

	pHash, err := app.hashAapter.FromMultiBytes([][]byte{
		[]byte(strings.Join(app.variables, ",")),
		app.operation.Hash().Bytes(),
		[]byte(isInitial),
	})

	if err != nil {
		return nil, err
	}

	return createAssignment(
		*pHash,
		app.variables,
		app.operation,
		app.isInitial,
	), nil
}
