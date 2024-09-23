package schemas

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/domain/scripts/specifics/schemas/connections"
	"github.com/steve-care-software/steve/domain/scripts/specifics/schemas/points"
)

type builder struct {
	hashAdapter hash.Adapter
	head        heads.Head
	points      points.Points
	connections connections.Connections
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		head:        nil,
		points:      nil,
		connections: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithHead adds an head to the builder
func (app *builder) WithHead(head heads.Head) Builder {
	app.head = head
	return app
}

// WithPoints adds a points to the builder
func (app *builder) WithPoints(points points.Points) Builder {
	app.points = points
	return app
}

// WithConnections adds a connections to the builder
func (app *builder) WithConnections(connections connections.Connections) Builder {
	app.connections = connections
	return app
}

// Now builds a new Schema instance
func (app *builder) Now() (Schema, error) {
	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build a Schema instance")
	}

	if app.points == nil {
		return nil, errors.New("the points is mandatory in order to build a Schema instance")
	}

	if app.connections == nil {
		return nil, errors.New("the connections is mandatory in order to build a Schema instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.head.Hash().Bytes(),
		app.points.Hash().Bytes(),
		app.connections.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createSchema(
		*pHash,
		app.head,
		app.points,
		app.connections,
	), nil
}
