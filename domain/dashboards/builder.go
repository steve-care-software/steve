package dashboards

import (
	"errors"

	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/dashboards/stencils"
)

type builder struct {
	root    hash.Hash
	visitor hash.Hash
	library stencils.Stencils
}

func createBuilder() Builder {
	out := builder{
		root:    nil,
		visitor: nil,
		library: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root hash.Hash) Builder {
	app.root = root
	return app
}

// WithVisitor adds a visitor to the builder
func (app *builder) WithVisitor(visitor hash.Hash) Builder {
	app.visitor = visitor
	return app
}

// WithLibrary adds a library to the builder
func (app *builder) WithLibrary(library stencils.Stencils) Builder {
	app.library = library
	return app
}

// Now builds a new Dashboard instance
func (app *builder) Now() (Dashboard, error) {
	if app.root == nil {
		return nil, errors.New("the root stencil hash is mandatory in order to build a Dashboard instance")
	}

	if app.visitor == nil {
		return nil, errors.New("the visitor stencil hash is mandatory in order to build a Dashboard instance")
	}

	if app.library == nil {
		return nil, errors.New("the library stencils are mandatory in order to build a Dashboard instance")
	}

	root, err := app.library.Fetch(app.root)
	if err != nil {
		return nil, err
	}

	visitor, err := app.library.Fetch(app.visitor)
	if err != nil {
		return nil, err
	}

	return createDashboard(root, visitor, app.library), nil
}
