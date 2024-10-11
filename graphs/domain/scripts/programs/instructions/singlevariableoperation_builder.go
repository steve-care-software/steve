package instructions

import (
	"errors"
)

type singleVariableOperationBuilder struct {
	name      string
	pOperator *uint8
}

func createSingleVariableOperationBuilder() SingleVariableOperationBuilder {
	return &singleVariableOperationBuilder{
		name:      "",
		pOperator: nil,
	}
}

// Create initializes the single variable operation builder
func (app *singleVariableOperationBuilder) Create() SingleVariableOperationBuilder {
	return createSingleVariableOperationBuilder()
}

// WithName adds a name to the single variable operation builder
func (app *singleVariableOperationBuilder) WithName(name string) SingleVariableOperationBuilder {
	app.name = name
	return app
}

// WithOperator adds an operator to the single variable operation builder
func (app *singleVariableOperationBuilder) WithOperator(operator uint8) SingleVariableOperationBuilder {
	app.pOperator = &operator
	return app
}

// Now builds a new SingleVariableOperation instance
func (app *singleVariableOperationBuilder) Now() (SingleVariableOperation, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory to build a SingleVariableOperation instance")
	}
	if app.pOperator == nil {
		return nil, errors.New("the operator is mandatory to build a SingleVariableOperation instance")
	}

	return createSingleVariableOperation(app.name, *app.pOperator), nil
}
