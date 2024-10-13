package conditions

import (
	"errors"
)

type clauseBuilder struct {
	pOperator *uint8
	element   Element
}

func createClauseBuilder() ClauseBuilder {
	return &clauseBuilder{
		pOperator: nil,
		element:   nil,
	}
}

// Create initializes the clause builder
func (app *clauseBuilder) Create() ClauseBuilder {
	return createClauseBuilder()
}

// WithOperator adds an operator to the clause builder
func (app *clauseBuilder) WithOperator(operator uint8) ClauseBuilder {
	app.pOperator = &operator
	return app
}

// WithElement adds an element to the clause builder
func (app *clauseBuilder) WithElement(element Element) ClauseBuilder {
	app.element = element
	return app
}

// Now builds a new Clause instance
func (app *clauseBuilder) Now() (Clause, error) {
	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Clause instance")
	}

	if app.pOperator == nil {
		return nil, errors.New("the operator is mandatory in order to build a Clause instance")
	}

	return createClause(*app.pOperator, app.element), nil
}
