package connections

import (
	"errors"

	"github.com/steve-care-software/steve/domain/connections"
	"github.com/steve-care-software/steve/domain/points"
)

type inMemoryBuilder struct {
	connectionsBuilder connections.Builder
	connections        connections.Connections
	points             points.Points
}

func createInMemoryBuilder(
	connectionsBuilder connections.Builder,
) InMemoryBuilder {
	out := inMemoryBuilder{
		connectionsBuilder: connectionsBuilder,
		connections:        nil,
		points:             nil,
	}

	return &out
}

// Create initializes the builder
func (app *inMemoryBuilder) Create() InMemoryBuilder {
	return createInMemoryBuilder(
		app.connectionsBuilder,
	)
}

// WithConnections add connections to the builder
func (app *inMemoryBuilder) WithConnections(connections connections.Connections) InMemoryBuilder {
	app.connections = connections
	return app
}

// WithPoints add points to the builder
func (app *inMemoryBuilder) WithPoints(points points.Points) InMemoryBuilder {
	app.points = points
	return app
}

// Now builds a new Application instance
func (app *inMemoryBuilder) Now() (Application, error) {
	if app.connections == nil {
		return nil, errors.New("the Connections is mandatory in order to build an Application instance")
	}

	mpFromList := map[string][]connections.Connection{}
	list := app.connections.List()
	for _, oneConnection := range list {
		keyname := oneConnection.From().String()
		if _, ok := mpFromList[keyname]; !ok {
			mpFromList[keyname] = []connections.Connection{}
		}

		mpFromList[keyname] = append(mpFromList[keyname], oneConnection)
	}

	mpFromConns := map[string]connections.Connections{}
	for keyname, list := range mpFromList {
		ins, err := app.connectionsBuilder.Create().WithList(list).Now()
		if err != nil {
			return nil, err
		}

		mpFromConns[keyname] = ins
	}

	return createApplication(
		mpFromConns,
	), nil
}
