package schemas

import (
	"errors"

	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/headers"
)

type builder struct {
	header      headers.Header
	points      []string
	connections connections.Connections
}

func createBuilder() Builder {
	out := builder{
		header:      nil,
		points:      nil,
		connections: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithHeader adds an header to the builder
func (app *builder) WithHeader(header headers.Header) Builder {
	app.header = header
	return app
}

// WithPoints add points to the builder
func (app *builder) WithPoints(points []string) Builder {
	app.points = points
	return app
}

// WithConnections add connections to the builder
func (app *builder) WithConnections(connections connections.Connections) Builder {
	app.connections = connections
	return app
}

// Now builds a new Schema instance
func (app *builder) Now() (Schema, error) {
	if app.header == nil {
		return nil, errors.New("the header is mandatory in order to build a Schema instance")
	}

	if app.points != nil && len(app.points) <= 0 {
		app.points = nil
	}

	if app.points == nil {
		return nil, errors.New("the points is mandatory in order to build a Schema instance")
	}

	if app.connections == nil {
		return nil, errors.New("the connections is mandatory in order to build a Schema instance")
	}

	return createSchema(
		app.header,
		app.points,
		app.connections,
	), nil
}
