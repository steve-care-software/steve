package points

import (
	"errors"

	"github.com/steve-care-software/steve/domain/points/bridges"
	"github.com/steve-care-software/steve/domain/points/contexts"
)

type pointBuilder struct {
	context contexts.Context
	bridge  bridges.Bridge
	from    []byte
}

func createPointBuilder() PointBuilder {
	out := pointBuilder{
		context: nil,
		bridge:  nil,
		from:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *pointBuilder) Create() PointBuilder {
	return createPointBuilder()
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

	return createPoint(app.context, app.bridge, app.from), nil
}
