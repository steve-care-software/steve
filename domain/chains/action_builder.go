package chains

import "errors"

type actionBuilder struct {
	interpreter Interpreter
	transpile   Transpile
}

func createActionBuilder() ActionBuilder {
	out := actionBuilder{
		interpreter: nil,
		transpile:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *actionBuilder) Create() ActionBuilder {
	return createActionBuilder()
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

	if app.interpreter != nil {
		return createActionWithInterpreter(app.interpreter), nil
	}

	if app.transpile != nil {
		return createActionWithTranspile(app.transpile), nil
	}

	return nil, errors.New("the Action is invalid")

}
