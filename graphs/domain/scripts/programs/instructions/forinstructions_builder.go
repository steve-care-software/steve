package instructions

import (
	"errors"
)

type forInstructionsBuilder struct {
	list []ForInstruction
}

func createForInstructionsBuilder() ForInstructionsBuilder {
	out := forInstructionsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *forInstructionsBuilder) Create() ForInstructionsBuilder {
	return createForInstructionsBuilder()
}

// WithList adds a list to the builder
func (app *forInstructionsBuilder) WithList(list []ForInstruction) ForInstructionsBuilder {
	app.list = list
	return app
}

// Now builds a new ForInstructions instance
func (app *forInstructionsBuilder) Now() (ForInstructions, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 ForInstruction in order to build a ForInstructions instance")
	}

	return createForInstructions(app.list), nil
}
