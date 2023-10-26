package preparations

import "errors"

type conditionBuilder struct {
	variable     string
	preparations Preparations
}

func createConditionBuilder() ConditionBuilder {
	out := conditionBuilder{
		variable:     "",
		preparations: nil,
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

// WithPreparations add preparations to the builder
func (app *conditionBuilder) WithPreparations(preparations Preparations) ConditionBuilder {
	app.preparations = preparations
	return app
}

// Now builds a new Condition instance
func (app *conditionBuilder) Now() (Condition, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a Condition instance")
	}

	if app.preparations == nil {
		return nil, errors.New("the preparations is mandatory in order to build a Condition instance")
	}

	return createCondition(app.variable, app.preparations), nil
}
