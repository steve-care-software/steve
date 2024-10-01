package links

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/schemas/points"
)

type linkBuilder struct {
	hashAdapter hash.Adapter
	origin      points.Point
	target      points.Point
}

func createLinkBuilder(
	hashAdapter hash.Adapter,
) LinkBuilder {
	out := linkBuilder{
		hashAdapter: hashAdapter,
		origin:      nil,
		target:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *linkBuilder) Create() LinkBuilder {
	return createLinkBuilder(
		app.hashAdapter,
	)
}

// WithOrigin adds an origin to the builder
func (app *linkBuilder) WithOrigin(origin points.Point) LinkBuilder {
	app.origin = origin
	return app
}

// WithTarget adds a target to the builder
func (app *linkBuilder) WithTarget(target points.Point) LinkBuilder {
	app.target = target
	return app
}

// Now builds a new Link instance
func (app *linkBuilder) Now() (Link, error) {
	if app.origin == nil {
		return nil, errors.New("the origin is mandatory in order to build a Link instance")
	}

	if app.target == nil {
		return nil, errors.New("the target is mandatory in order to build a Link instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.origin.Hash().Bytes(),
		app.target.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createLink(*pHash, app.origin, app.target), nil
}
