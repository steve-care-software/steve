package saves

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/assignments"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/conditions"
)

type builder struct {
	assignment assignments.Assignment
	condition  conditions.Condition
}

func createBuilder() Builder {
	return &builder{
		assignment: nil,
		condition:  nil,
	}
}

// Create initializes the save builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithAssignment adds an assignment to the save builder
func (app *builder) WithAssignment(assignment assignments.Assignment) Builder {
	app.assignment = assignment
	return app
}

// WithCondition adds a condition to the save builder
func (app *builder) WithCondition(condition conditions.Condition) Builder {
	app.condition = condition
	return app
}

// Now builds a new Save instance
func (app *builder) Now() (Save, error) {
	if app.assignment == nil {
		return nil, errors.New("the assignment is mandatory in order to build a Save instance")
	}

	if app.condition != nil {
		return createSaveWithCondition(app.assignment, app.condition), nil
	}

	return createSave(app.assignment), nil
}
