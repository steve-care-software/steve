package assignments

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/containers/kinds"
)

type assigneeBuilder struct {
	name AssigneeName
	kind kinds.Kind
}

func createAssigneeBuilder() AssigneeBuilder {
	return &assigneeBuilder{
		name: nil,
		kind: nil,
	}
}

// Create initializes the AssigneeBuilder
func (app *assigneeBuilder) Create() AssigneeBuilder {
	return createAssigneeBuilder()
}

// WithName sets the name field in the builder
func (app *assigneeBuilder) WithName(name AssigneeName) AssigneeBuilder {
	app.name = name
	return app
}

// WithKind sets the kind field in the builder
func (app *assigneeBuilder) WithKind(kind kinds.Kind) AssigneeBuilder {
	app.kind = kind
	return app
}

// Now builds a new Assignee instance
func (app *assigneeBuilder) Now() (Assignee, error) {
	if app.name == nil {
		return nil, errors.New("the name field is mandatory to build an Assignee instance")
	}

	if app.kind != nil {
		return createAssigneeWithKind(app.name, app.kind), nil
	}

	return createAssignee(app.name), nil
}
