package conditions

import (
	"errors"
)

type builder struct {
	element Element
	clauses Clauses
}

func createBuilder() Builder {
	out := builder{
		element: nil,
		clauses: nil,
	}

	return &out
}

// Create initializes the condition builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithElement adds an element to the condition builder
func (app *builder) WithElement(element Element) Builder {
	app.element = element
	return app
}

// WithClauses adds clauses to the condition builder
func (app *builder) WithClauses(clauses Clauses) Builder {
	app.clauses = clauses
	return app
}

// Now builds a new Condition instance
func (app *builder) Now() (Condition, error) {
	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Condition instance")
	}

	if app.clauses != nil {
		return createConditionWithClauses(
			app.element,
			app.clauses,
		), nil
	}

	return createCondition(
		app.element,
	), nil
}
