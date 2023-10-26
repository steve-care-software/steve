package symbols

import "errors"

type builder struct {
	list []Symbol
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Symbol) Builder {
	app.list = list
	return app
}

// Now builds a new Symbols instance
func (app *builder) Now() (Symbols, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Symbol in order to build a Symbols instance")
	}

	return createSymbols(app.list), nil
}
