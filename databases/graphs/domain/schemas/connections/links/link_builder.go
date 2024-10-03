package links

import (
	"errors"

	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/links/references"
)

type linkBuilder struct {
	origin references.Reference
	target references.Reference
}

func createLinkBuilder() LinkBuilder {
	out := linkBuilder{
		origin: nil,
		target: nil,
	}

	return &out
}

// Create initializes the builder
func (app *linkBuilder) Create() LinkBuilder {
	return createLinkBuilder()
}

// WithOrigin adds an origin to the builder
func (app *linkBuilder) WithOrigin(origin references.Reference) LinkBuilder {
	app.origin = origin
	return app
}

// WithTarget adds a target to the builder
func (app *linkBuilder) WithTarget(target references.Reference) LinkBuilder {
	app.target = target
	return app
}

// Now builds a new Link instance
func (app *linkBuilder) Now() (Link, error) {
	if app.origin == nil {
		return nil, errors.New("the origin is mandatory in order to build a Link instance")
	}

	if app.target == nil {
		return nil, errors.New("the target is mandatory in order to build a Link instance")
	}

	return createLink(
		app.origin,
		app.target,
	), nil
}
