package symbols

import "errors"

type builder struct {
	name  string
	pKind *uint8
}

func createBuilder() Builder {
	out := builder{
		name:  "",
		pKind: nil,
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
func (app *builder) WithKind(kind uint8) Builder {
	app.pKind = &kind
	return app
}

// Now builds a new Symbol instance
func (app *builder) Now() (Symbol, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Symbol instance")
	}

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Symbol instance")
	}

	kind := *app.pKind
	if kind&KindBytes|KindLayer|KindLink == 0 {
		return nil, errors.New("the kind is invalid when creating a Symbol instance")
	}

	return createSymbol(app.name, kind), nil
}
