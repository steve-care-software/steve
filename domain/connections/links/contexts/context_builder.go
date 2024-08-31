package contexts

import (
	"errors"

	"github.com/google/uuid"
)

type contextBuilder struct {
	pIdentifier *uuid.UUID
	name        string
	pParent     *uuid.UUID
}

func createContextBuilder() ContextBuilder {
	out := contextBuilder{
		pIdentifier: nil,
		name:        "",
		pParent:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *contextBuilder) Create() ContextBuilder {
	return createContextBuilder()
}

// WithIdentifier adds an identifier to the builder
func (app *contextBuilder) WithIdentifier(identifier uuid.UUID) ContextBuilder {
	app.pIdentifier = &identifier
	return app
}

// WithName adds a name to the builder
func (app *contextBuilder) WithName(name string) ContextBuilder {
	app.name = name
	return app
}

// WithParent adds a parent to the builder
func (app *contextBuilder) WithParent(parent uuid.UUID) ContextBuilder {
	app.pParent = &parent
	return app
}

// Now builds a new Context instance
func (app *contextBuilder) Now() (Context, error) {
	if app.pIdentifier == nil {
		return nil, errors.New("the identifier is mandatory in order to build a Context instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Context instance")
	}

	if app.pParent != nil {
		return createContextWithParent(*app.pIdentifier, app.name, app.pParent), nil
	}

	return createContext(*app.pIdentifier, app.name), nil
}
