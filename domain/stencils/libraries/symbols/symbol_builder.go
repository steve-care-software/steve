package symbols

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links"
)

type symbolBuilder struct {
	hashAdapter hash.Adapter
	bytes       []byte
	layer       layers.Layer
	link        links.Link
}

func createSymbolBuilder(
	hashAdapter hash.Adapter,
) SymbolBuilder {
	out := symbolBuilder{
		hashAdapter: hashAdapter,
		bytes:       nil,
		layer:       nil,
		link:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *symbolBuilder) Create() SymbolBuilder {
	return createSymbolBuilder(
		app.hashAdapter,
	)
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

	data := [][]byte{}
	if app.bytes != nil {
		data = append(data, app.bytes)
	}

	if app.layer != nil {
		data = append(data, app.layer.Hash().Bytes())
	}

	if app.link != nil {
		data = append(data, app.link.Hash().Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the Symbol is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.bytes != nil {
		return createSymbolWithBytes(*pHash, app.bytes), nil
	}

	if app.layer != nil {
		return createSymbolWithLayer(*pHash, app.layer), nil
	}

	return createSymbolWithLink(*pHash, app.link), nil
}
