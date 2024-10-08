package bridges

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/steve/engine/domain/graphs/connections"
	"github.com/steve-care-software/steve/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	connection  connections.Connection
	weight      float32
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		connection:  nil,
		weight:      0.0,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithConnection adds a connection to the builder
func (app *builder) WithConnection(connection connections.Connection) Builder {
	app.connection = connection
	return app
}

// WithWeight adds a weight to the builder
func (app *builder) WithWeight(weight float32) Builder {
	app.weight = weight
	return app
}

// Now builds a new Bridge instance
func (app *builder) Now() (Bridge, error) {
	if app.connection == nil {
		return nil, errors.New("the connection is mandatory in order to build a Bridge instance")
	}

	data := [][]byte{
		app.connection.Hash().Bytes(),
	}

	if app.weight > 0.0 {
		data = append(data, []byte(strconv.FormatFloat(float64(app.weight), 'f', 10, 32)))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.weight > 0.0 {
		return createBridgeWithWeight(*pHash, app.connection, app.weight), nil
	}

	return createBridge(*pHash, app.connection), nil
}
