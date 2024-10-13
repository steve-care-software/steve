package links

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/assignments"
)

type builder struct {
	origin assignments.Assignment
	target assignments.Assignment
}

func createBuilder() Builder {
	out := builder{
		origin: nil,
		target: nil,
	}

	return &out
}

// Create initializes the link builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithOrigin adds an origin assignment to the link builder
func (app *builder) WithOrigin(origin assignments.Assignment) Builder {
	app.origin = origin
	return app
}

// WithTarget adds a target assignment to the link builder
func (app *builder) WithTarget(target assignments.Assignment) Builder {
	app.target = target
	return app
}

// Now builds a new Link instance
func (app *builder) Now() (Link, error) {
	if app.origin == nil {
		return nil, errors.New("the origin is mandatory in order to build a Link instance")
	}

	if app.target == nil {
		return nil, errors.New("the target is mandatory in order to build a Link instance")
	}

	return createLink(app.origin, app.target), nil
}
