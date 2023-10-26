package expectations

import (
	"errors"

	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns/kinds"
)

type builder struct {
	variable string
	kind     kinds.Kind
}

func createBuilder() Builder {
	out := builder{
		variable: "",
		kind:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithVariable adds a variable to the builder
func (app *builder) WithVariable(variable string) Builder {
	app.variable = app.variable
	return app
}

// WithKind adds a kind to the builder
func (app *builder) WithKind(kind kinds.Kind) Builder {
	app.kind = app.kind
	return app
}

// Now builds a new Expectation instance
func (app *builder) Now() (Expectation, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build an Expectation instance")
	}

	if app.kind == nil {
		return nil, errors.New("the kind is mandatory in order to build an Expectation instance")
	}

	return createExpectation(app.variable, app.kind), nil
}
