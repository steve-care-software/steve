package symbols

import (
	"errors"

	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links"
)

type symbolBuilder struct {
	bytes []byte
	layer layers.Layer
	link  links.Link
}

func createSymbolBuilder() SymbolBuilder {
	out := symbolBuilder{
		bytes: nil,
		layer: nil,
		link:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *symbolBuilder) Create() SymbolBuilder {
	return createSymbolBuilder()
}

// WithBytes add bytes to the builder
func (app *symbolBuilder) WithBytes(bytes []byte) SymbolBuilder {
	app.bytes = bytes
	return app
}

// WithLayer adds a layer to the builder
func (app *symbolBuilder) WithLayer(layer layers.Layer) SymbolBuilder {
	app.layer = layer
	return app
}

// WithLink adds a link to the builder
func (app *symbolBuilder) WithLink(link links.Link) SymbolBuilder {
	app.link = link
	return app
}

// Now builds a new Symbol instance
func (app *symbolBuilder) Now() (Symbol, error) {
	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.bytes != nil {
		return createSymbolWithBytes(app.bytes), nil
	}

	if app.layer != nil {
		return createSymbolWithLayer(app.layer), nil
	}

	if app.link != nil {
		return createSymbolWithLink(app.link), nil
	}

	return nil, errors.New("the Symbol is invalid")
}
