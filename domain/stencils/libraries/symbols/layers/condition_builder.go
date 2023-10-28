package layers

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type conditionBuilder struct {
	hashAdapter hash.Adapter
	variable    string
	executions  Executions
}

func createConditionBuilder(
	hashAdapter hash.Adapter,
) ConditionBuilder {
	out := conditionBuilder{
		hashAdapter: hashAdapter,
		variable:    "",
		executions:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *conditionBuilder) Create() ConditionBuilder {
	return createConditionBuilder(
		app.hashAdapter,
	)
}

// WithVariable adds a variable to the builder
func (app *conditionBuilder) WithVariable(variable string) ConditionBuilder {
	app.variable = variable
	return app
}

// WithExecutions adds an executions to the builder
func (app *conditionBuilder) WithExecutions(executions Executions) ConditionBuilder {
	app.executions = executions
	return app
}

// Now builds a new Condition instance
func (app *conditionBuilder) Now() (Condition, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a Condition instance")
	}

	if app.executions == nil {
		return nil, errors.New("the executions is mandatory in order to build a Condition instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.variable),
		app.executions.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createCondition(*pHash, app.variable, app.executions), nil
}
