package points

import (
	"errors"

	"github.com/steve-care-software/steve/domain/points/bridges"
)

type pointBuilder struct {
	bridge bridges.Bridge
	from   []byte
}

func createPointBuilder() PointBuilder {
	out := pointBuilder{
		bridge: nil,
		from:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *pointBuilder) Create() PointBuilder {
	return createPointBuilder()
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
	if app.bridge == nil {
		return nil, errors.New("the bridge is mandatory in order to build a Point instance")
	}

	if app.from != nil && len(app.from) <= 0 {
		app.from = nil
	}

	if app.from == nil {
		return nil, errors.New("the from data is mandatory in order to build a Point instance")
	}

	return createPoint(app.bridge, app.from), nil
}
