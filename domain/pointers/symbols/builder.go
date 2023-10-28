package symbols

import (
	"errors"

	"github.com/steve-care-software/steve/domain/pointers/symbols/kinds"
)

type builder struct {
	name string
	kind kinds.Kind
}

func createBuilder() Builder {
	out := builder{
		name: "",
		kind: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithKind adds a kind to the builder
func (app *builder) WithKind(kind kinds.Kind) Builder {
	app.kind = kind
	return app
}

// Now builds a new Symbol instance
func (app *builder) Now() (Symbol, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Symbol instance")
	}

	if app.kind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Symbol instance")
	}

	return createSymbol(app.name, app.kind), nil
}
