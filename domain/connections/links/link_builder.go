package links

import (
	"errors"

	"github.com/steve-care-software/steve/domain/connections/links/contexts"
)

type linkBuilder struct {
	name     string
	isLeft   bool
	weight   float32
	contexts contexts.Contexts
}

func createLinkBuilder() LinkBuilder {
	out := linkBuilder{
		name:     "",
		isLeft:   false,
		weight:   0.0,
		contexts: nil,
	}

	return &out
}

// Create initializes the builder
func (app *linkBuilder) Create() LinkBuilder {
	return createLinkBuilder()
}

// WithContexts add contexts to the builder
func (app *linkBuilder) WithContexts(contexts contexts.Contexts) LinkBuilder {
	app.contexts = contexts
	return app
}

// WithName adds a name to the builder
func (app *linkBuilder) WithName(name string) LinkBuilder {
	app.name = name
	return app
}

// WithWeight adds a weight to the builder
func (app *linkBuilder) WithWeight(weight float32) LinkBuilder {
	app.weight = weight
	return app
}

// IsLeft flags the builder as left
func (app *linkBuilder) IsLeft() LinkBuilder {
	app.isLeft = true
	return app
}

// Now builds a new Link instance
func (app *linkBuilder) Now() (Link, error) {
	if app.contexts == nil {
		return nil, errors.New("the contexts is mandatory in order to build a Link instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Link instance")
	}

	if app.weight <= 0.0 {
		return nil, errors.New("the weight must be greater than 0.0 in order to build a Link instance")
	}

	return createLink(app.name, app.isLeft, app.weight, app.contexts), nil
}
