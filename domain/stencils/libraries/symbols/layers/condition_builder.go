package layers

import "errors"

type conditionBuilder struct {
	variable   string
	executions Executions
}

func createConditionBuilder() ConditionBuilder {
	out := conditionBuilder{
		variable:   "",
		executions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *conditionBuilder) Create() ConditionBuilder {
	return createConditionBuilder()
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

	return createCondition(app.variable, app.executions), nil
}
