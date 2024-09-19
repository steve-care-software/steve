package chains

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type interpreterBuilder struct {
	hashAdapter hash.Adapter
	variable    string
	next        Chain
}

func createInterpreterBuilder(
	hashAdapter hash.Adapter,
) InterpreterBuilder {
	out := interpreterBuilder{
		hashAdapter: hashAdapter,
		variable:    "",
		next:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *interpreterBuilder) Create() InterpreterBuilder {
	return createInterpreterBuilder(
		app.hashAdapter,
	)
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

	data := [][]byte{
		[]byte(app.variable),
	}

	if app.next != nil {
		data = append(data, app.next.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.next != nil {
		return createInterpreterWithNext(*pHash, app.variable, app.next), nil
	}

	return createInterpreter(*pHash, app.variable), nil
}
