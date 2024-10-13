package assignments

import (
	"errors"
)

type assigneesBuilder struct {
	list []Assignee
}

func createAssigneesBuilder() AssigneesBuilder {
	out := assigneesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *assigneesBuilder) Create() AssigneesBuilder {
	return createAssigneesBuilder()
}

// WithList adds a list to the builder
func (app *assigneesBuilder) WithList(list []Assignee) AssigneesBuilder {
	app.list = list
	return app
}

// Now builds a new Assignees instance
func (app *assigneesBuilder) Now() (Assignees, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Assignee in order to build a Assignees instance")
	}

	return createAssignees(app.list), nil
}
