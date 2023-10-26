package layers

import "errors"

type valueAssignmentsBuilder struct {
	list []ValueAssignment
}

func createValueAssignmentsBuilder() ValueAssignmentsBuilder {
	out := valueAssignmentsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the valueAssignmentsBuilder
func (app *valueAssignmentsBuilder) Create() ValueAssignmentsBuilder {
	return createValueAssignmentsBuilder()
}

// WithList adds a list to the valueAssignmentsBuilder
func (app *valueAssignmentsBuilder) WithList(list []ValueAssignment) ValueAssignmentsBuilder {
	app.list = list
	return app
}

// Now builds a new ValueAssignments instance
func (app *valueAssignmentsBuilder) Now() (ValueAssignments, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 ValueAssignment in order to build a ValueAssignments instance")
	}

	return createValueAssignments(app.list), nil
}
