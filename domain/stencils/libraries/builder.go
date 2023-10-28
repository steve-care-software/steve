package libraries

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols"
)

type builder struct {
	hashAdapter hash.Adapter
	path        []string
	symbols     symbols.Symbols
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		path:        nil,
		symbols:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithPath adds a path to the builder
func (app *builder) WithPath(path []string) Builder {
	app.path = path
	return app
}

// WithSybols add symbols to the builder
func (app *builder) WithSybols(symbols symbols.Symbols) Builder {
	app.symbols = symbols
	return app
}

// Now builds a new Library instance
func (app *builder) Now() (Library, error) {
	if app.path == nil {
		app.path = []string{}
	}

	if app.symbols == nil {
		return nil, errors.New("the symbols are mandatory in order to build a Library instance")
	}

	data := [][]byte{
		app.symbols.Hash().Bytes(),
	}

	for _, oneFolder := range app.path {
		data = append(data, []byte(oneFolder))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createLibrary(*pHash, app.path, app.symbols), nil
}
