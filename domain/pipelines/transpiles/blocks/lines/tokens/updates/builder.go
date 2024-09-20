package updates

import (
	"errors"

	"github.com/steve-care-software/steve/domain/pipelines/transpiles/blocks/lines/tokens/pointers"
	"github.com/steve-care-software/steve/domain/pipelines/transpiles/blocks/lines/tokens/updates/targets"
)

type builder struct {
	origin pointers.Pointer
	target targets.Target
}

func createBuilder() Builder {
	out := builder{
		origin: nil,
		target: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithOrigin adds an origin to the builder
func (app *builder) WithOrigin(origin pointers.Pointer) Builder {
	app.origin = origin
	return app
}

// WithTarget adds a target to the builder
func (app *builder) WithTarget(target targets.Target) Builder {
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

	return createUpdate(app.origin, app.target), nil
}
