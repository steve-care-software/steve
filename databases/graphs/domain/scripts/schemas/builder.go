package schemas

import (
	"errors"

	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections"
)

type builder struct {
	head        heads.Head
	points      []string
	connections connections.Connections
}

func createBuilder() Builder {
	out := builder{
		head:        nil,
		points:      nil,
		connections: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithHead adds an header to the builder
func (app *builder) WithHead(head heads.Head) Builder {
	app.head = head
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
	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build a Schema instance")
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
		app.head,
		app.points,
		app.connections,
	), nil
}
