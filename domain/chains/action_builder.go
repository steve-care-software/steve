package chains

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type actionBuilder struct {
	hashAdapter hash.Adapter
	interpreter Interpreter
	transpile   Transpile
}

func createActionBuilder(
	hashAdapter hash.Adapter,
) ActionBuilder {
	out := actionBuilder{
		hashAdapter: hashAdapter,
		interpreter: nil,
		transpile:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *actionBuilder) Create() ActionBuilder {
	return createActionBuilder(
		app.hashAdapter,
	)
}

// WithInterpret adds an interpreter to the builder
func (app *actionBuilder) WithInterpret(interpret Interpreter) ActionBuilder {
	app.interpreter = interpret
	return app
}

// WithTranspile adds a transpile to the builder
func (app *actionBuilder) WithTranspile(transpile Transpile) ActionBuilder {
	app.transpile = transpile
	return app
}

// Now builds a new Action instance
func (app *actionBuilder) Now() (Action, error) {
	data := [][]byte{}
	if app.interpreter != nil {
		data = append(data, app.interpreter.Hash().Bytes())
	}

	if app.transpile != nil {
		data = append(data, app.transpile.Hash().Bytes())
	}

	if len(data) != 1 {
		return nil, errors.New("the Action is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.interpreter != nil {
		return createActionWithInterpreter(*pHash, app.interpreter), nil
	}

	return createActionWithTranspile(*pHash, app.transpile), nil

}
