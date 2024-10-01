package contexts

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/hash"
)

type contextBuilder struct {
	hashAdapter hash.Adapter
	name        string
	parent      hash.Hash
}

func createContextBuilder(
	hashAdapter hash.Adapter,
) ContextBuilder {
	out := contextBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		parent:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *contextBuilder) Create() ContextBuilder {
	return createContextBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *contextBuilder) WithName(name string) ContextBuilder {
	app.name = name
	return app
}

// WithParent adds a parent to the builder
func (app *contextBuilder) WithParent(parent hash.Hash) ContextBuilder {
	app.parent = parent
	return app
}

// Now builds a new Context instance
func (app *contextBuilder) Now() (Context, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Context instance")
	}

	data := [][]byte{
		[]byte(app.name),
	}

	if app.parent != nil {
		data = append(data, app.parent.Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.parent != nil {
		return createContextWithParent(*pHash, app.name, app.parent), nil
	}

	return createContext(*pHash, app.name), nil
}
