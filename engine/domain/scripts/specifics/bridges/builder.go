package bridges

import (
	"errors"

	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/bridges/connections"
)

type builder struct {
	hashAdapter hash.Adapter
	head        heads.Head
	origin      string
	target      string
	connections connections.Connections
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		head:        nil,
		origin:      "",
		target:      "",
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

// WithHead adds a head to the builder
func (app *builder) WithHead(head heads.Head) Builder {
	app.head = head
	return app
}

// WithOrigin adds an origin to the builder
func (app *builder) WithOrigin(origin string) Builder {
	app.origin = origin
	return app
}

// WithTarget adds a target to the builder
func (app *builder) WithTarget(target string) Builder {
	app.target = target
	return app
}

// WithConnections add connections to the builder
func (app *builder) WithConnections(connections connections.Connections) Builder {
	app.connections = connections
	return app
}

// Now builds a new Bridge instance
func (app *builder) Now() (Bridge, error) {
	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build a Bridge instance")
	}

	if app.origin == "" {
		return nil, errors.New("the origin is mandatory in order to build a Bridge instance")
	}

	if app.target == "" {
		return nil, errors.New("the target is mandatory in order to build a Bridge instance")
	}

	if app.connections == nil {
		return nil, errors.New("the connections is mandatory in order to build a Bridge instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.head.Hash().Bytes(),
		[]byte(app.origin),
		[]byte(app.target),
		app.connections.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createBridge(
		*pHash,
		app.head,
		app.origin,
		app.target,
		app.connections,
	), nil
}

/*

head        heads.Head
	origin      string
	target      string
	connections connections.Connections

*/
