package points

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/graphs/points/bridges"
	"github.com/steve-care-software/steve/engine/domain/graphs/points/contexts"
	"github.com/steve-care-software/steve/engine/domain/hash"
)

type pointBuilder struct {
	hashAdapter hash.Adapter
	context     contexts.Context
	bridge      bridges.Bridge
	from        []byte
}

func createPointBuilder(
	hashAdapter hash.Adapter,
) PointBuilder {
	out := pointBuilder{
		hashAdapter: hashAdapter,
		context:     nil,
		bridge:      nil,
		from:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *pointBuilder) Create() PointBuilder {
	return createPointBuilder(
		app.hashAdapter,
	)
}

// WithContext adds a context to the builder
func (app *pointBuilder) WithContext(context contexts.Context) PointBuilder {
	app.context = context
	return app
}

// WithBridge adds a bridge to the builder
func (app *pointBuilder) WithBridge(bridge bridges.Bridge) PointBuilder {
	app.bridge = bridge
	return app
}

// From adds the from data to the builder
func (app *pointBuilder) From(from []byte) PointBuilder {
	app.from = from
	return app
}

// Now builds a new Point instance
func (app *pointBuilder) Now() (Point, error) {
	if app.context == nil {
		return nil, errors.New("the context is mandatory in order to build a Point instance")
	}

	if app.bridge == nil {
		return nil, errors.New("the bridge is mandatory in order to build a Point instance")
	}

	if app.from != nil && len(app.from) <= 0 {
		app.from = nil
	}

	if app.from == nil {
		return nil, errors.New("the from data is mandatory in order to build a Point instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.context.Hash().Bytes(),
		app.bridge.Hash().Bytes(),
		[]byte(app.from),
	})

	if err != nil {
		return nil, err
	}

	return createPoint(*pHash, app.context, app.bridge, app.from), nil
}
