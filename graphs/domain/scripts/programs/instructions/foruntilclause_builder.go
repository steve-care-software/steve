package instructions

import "errors"

type forUntilClauseBuilder struct {
	name   string
	pValue *uint
}

func createForUntilClauseBuilder() ForUntilClauseBuilder {
	return &forUntilClauseBuilder{
		name:   "",
		pValue: nil,
	}
}

// Create initializes the builder
func (app *forUntilClauseBuilder) Create() ForUntilClauseBuilder {
	return createForUntilClauseBuilder()
}

// WithName adds a name to the builder
func (app *forUntilClauseBuilder) WithName(name string) ForUntilClauseBuilder {
	app.name = name
	return app
}

// WithValue adds a value to the builder
func (app *forUntilClauseBuilder) WithValue(value uint) ForUntilClauseBuilder {
	app.pValue = &value
	return app
}

// Now builds and returns a ForUntilClause instance
func (app *forUntilClauseBuilder) Now() (ForUntilClause, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory to build a ForUntilClause instance")
	}

	if app.pValue == nil {
		return nil, errors.New("the value is mandatory to build a ForUntilClause instance")
	}

	return createForUntilClause(app.name, *app.pValue), nil
}
