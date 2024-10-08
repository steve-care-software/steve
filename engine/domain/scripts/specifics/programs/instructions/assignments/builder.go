package assignments

import (
	"errors"
	"strings"

	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/containers"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/assignments/assignables"
	"github.com/steve-care-software/steve/hash"
)

type builder struct {
	hashAapter hash.Adapter
	variables  []string
	assignable assignables.Assignable
	initial    containers.Container
}

func createBuilder(
	hashAapter hash.Adapter,
) Builder {
	out := builder{
		hashAapter: hashAapter,
		variables:  nil,
		assignable: nil,
		initial:    nil,
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

// WithAssignable add opration to the builder
func (app *builder) WithAssignable(assignable assignables.Assignable) Builder {
	app.assignable = assignable
	return app
}

// WithInitial adds an initial to the builder
func (app *builder) WithInitial(initial containers.Container) Builder {
	app.initial = initial
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

	if app.assignable == nil {
		return nil, errors.New("the assignable is mandatory in order to build an Assignment instance")
	}

	data := [][]byte{
		[]byte(strings.Join(app.variables, ",")),
		app.assignable.Hash().Bytes(),
	}

	if app.initial != nil {
		data = append(data, app.initial.Hash().Bytes())
	}

	pHash, err := app.hashAapter.FromMultiBytes(data)

	if err != nil {
		return nil, err
	}

	if app.initial != nil {
		return createAssignmentWithInitial(
			*pHash,
			app.variables,
			app.assignable,
			app.initial,
		), nil
	}

	return createAssignment(
		*pHash,
		app.variables,
		app.assignable,
	), nil
}
