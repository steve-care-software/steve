package conditions

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/assignments"
)

// elementBuilder represents the implementation of the ElementBuilder interface
type elementBuilder struct {
	assignment assignments.Assignment
	condition  Condition
}

func createElementBuilder() ElementBuilder {
	return &elementBuilder{
		assignment: nil,
		condition:  nil,
	}
}

// Create initializes the element builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithAssignment adds an assignment to the element builder
func (app *elementBuilder) WithAssignment(assignment assignments.Assignment) ElementBuilder {
	app.assignment = assignment
	return app
}

// WithCondition adds a condition to the element builder
func (app *elementBuilder) WithCondition(condition Condition) ElementBuilder {
	app.condition = condition
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.assignment == nil {
		return nil, errors.New("the assignment is mandatory in order to build an Element instance")
	}

	if app.condition == nil {
		return nil, errors.New("the condition is mandatory in order to build an Element instance")
	}

	return createElement(app.assignment, app.condition), nil
}
