package points

import (
	"errors"

	"github.com/steve-care-software/steve/domain/connections"
)

type pointBuilder struct {
	connection connections.Connection
	from       []byte
}

func createPointBuilder() PointBuilder {
	out := pointBuilder{
		connection: nil,
		from:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *pointBuilder) Create() PointBuilder {
	return createPointBuilder()
}

// WithConnection adds a connection to the builder
func (app *pointBuilder) WithConnection(connection connections.Connection) PointBuilder {
	app.connection = connection
	return app
}

// From adds the from data to the builder
func (app *pointBuilder) From(from []byte) PointBuilder {
	app.from = from
	return app
}

// Now builds a new Point instance
func (app *pointBuilder) Now() (Point, error) {
	if app.connection == nil {
		return nil, errors.New("the connection is mandatory in order to build a Point instance")
	}

	if app.from != nil && len(app.from) <= 0 {
		app.from = nil
	}

	if app.from == nil {
		return nil, errors.New("the from data is mandatory in order to build a Point instance")
	}

	return createPoint(app.connection, app.from), nil
}
