package bridges

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/bridges/links"
)

type builder struct {
	weight uint
	origin links.Link
	target links.Link
}

func createBuilder() Builder {
	return &builder{
		weight: 0,
		origin: nil,
		target: nil,
	}
}

// Create initializes the bridge builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithWeight adds a weight to the bridge builder
func (app *builder) WithWeight(weight uint) Builder {
	app.weight = weight
	return app
}

// WithOrigin adds an origin link to the bridge builder
func (app *builder) WithOrigin(origin links.Link) Builder {
	app.origin = origin
	return app
}

// WithTarget adds a target link to the bridge builder
func (app *builder) WithTarget(target links.Link) Builder {
	app.target = target
	return app
}

// Now builds a new Bridge instance
func (app *builder) Now() (Bridge, error) {
	if app.weight <= 0 {
		return nil, errors.New("the weight must be greater than zero (0) in order to build a Bridge instance")
	}

	if app.origin == nil {
		return nil, errors.New("the origin is mandatory in order to build a Bridge instance")
	}

	if app.target == nil {
		return nil, errors.New("the target is mandatory in order to build a Bridge instance")
	}

	return createBridge(app.weight, app.origin, app.target), nil
}
