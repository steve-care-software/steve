package links

import (
	"errors"

	"github.com/steve-care-software/steve/domain/relations/data/connections/links/contexts"
)

type linkBuilder struct {
	name     string
	isLeft   bool
	contexts contexts.Contexts
}

func createLinkBuilder() LinkBuilder {
	out := linkBuilder{
		name:     "",
		isLeft:   false,
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

	return createLink(app.name, app.isLeft, app.contexts), nil
}
