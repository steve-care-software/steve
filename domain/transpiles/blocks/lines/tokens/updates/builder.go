package updates

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens/pointers"
)

type builder struct {
	hashAdapter hash.Adapter
	origin      pointers.Pointer
	target      pointers.Pointer
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		origin:      nil,
		target:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithOrigin adds an origin to the builder
func (app *builder) WithOrigin(origin pointers.Pointer) Builder {
	app.origin = origin
	return app
}

// WithTarget adds a target to the builder
func (app *builder) WithTarget(target pointers.Pointer) Builder {
	app.target = target
	return app
}

// Now builds a new Update instance
func (app *builder) Now() (Update, error) {
	if app.origin == nil {
		return nil, errors.New("the origin is mandatory in order to build an Update instance")
	}

	if app.target == nil {
		return nil, errors.New("the target is mandatory in order to build an Update instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.origin.Hash().Bytes(),
		app.target.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createUpdate(*pHash, app.origin, app.target), nil
}
