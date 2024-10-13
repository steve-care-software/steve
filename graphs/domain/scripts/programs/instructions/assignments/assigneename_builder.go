package assignments

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references"
)

type assigneeNameBuilder struct {
	references references.References
	index      *uint
}

func createAssigneeNameBuilder() AssigneeNameBuilder {
	return &assigneeNameBuilder{
		references: nil,
		index:      nil,
	}
}

// Create initializes the AssigneeNameBuilder
func (app *assigneeNameBuilder) Create() AssigneeNameBuilder {
	return createAssigneeNameBuilder()
}

// WithReferences adds references to the builder
func (app *assigneeNameBuilder) WithReferences(references references.References) AssigneeNameBuilder {
	app.references = references
	return app
}

// WithIndex adds an index to the builder
func (app *assigneeNameBuilder) WithIndex(index uint) AssigneeNameBuilder {
	app.index = &index
	return app
}

// Now builds a new AssigneeName instance
func (app *assigneeNameBuilder) Now() (AssigneeName, error) {
	if app.references == nil {
		return nil, errors.New("the references field is mandatory to build an AssigneeName instance")
	}

	if app.index != nil {
		return createAssigneeNameWithIndex(app.references, *app.index), nil
	}

	return createAssigneeName(app.references), nil
}
