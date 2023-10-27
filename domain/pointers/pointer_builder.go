package pointers

import (
	"errors"

	"github.com/steve-care-software/steve/domain/pointers/symbols"
)

type pointerBuilder struct {
	path   []string
	symbol symbols.Symbol
}

func createPointerBuilder() PointerBuilder {
	out := pointerBuilder{
		path:   nil,
		symbol: nil,
	}

	return &out
}

// Create initializes the builder
func (app *pointerBuilder) Create() PointerBuilder {
	return createPointerBuilder()
}

// WithPath adds a path to the builder
func (app *pointerBuilder) WithPath(path []string) PointerBuilder {
	app.path = path
	return app
}

// WithSymbol adds a symbol to the builder
func (app *pointerBuilder) WithSymbol(symbol symbols.Symbol) PointerBuilder {
	app.symbol = symbol
	return app
}

// Now builds a new Pointer instance
func (app *pointerBuilder) Now() (Pointer, error) {
	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("the path is mandatory in order to build a Pointer instance")
	}

	return createPointer(app.path, app.symbol), nil
}
