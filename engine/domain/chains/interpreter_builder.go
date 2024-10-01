package chains

import (
	"errors"
)

type interpreterBuilder struct {
	variable string
	next     Chain
}

func createInterpreterBuilder() InterpreterBuilder {
	out := interpreterBuilder{
		variable: "",
		next:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *interpreterBuilder) Create() InterpreterBuilder {
	return createInterpreterBuilder()
}

// WithVariable adds a variable to the builder
func (app *interpreterBuilder) WithVariable(variable string) InterpreterBuilder {
	app.variable = variable
	return app
}

// WithNext adds a next chain to the builder
func (app *interpreterBuilder) WithNext(next Chain) InterpreterBuilder {
	app.next = next
	return app
}

// Now builds
func (app *interpreterBuilder) Now() (Interpreter, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build an Interpreter")
	}

	if app.next != nil {
		return createInterpreterWithNext(app.variable, app.next), nil
	}

	return createInterpreter(app.variable), nil
}
