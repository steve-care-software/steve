package connections

import (
	"errors"

	"github.com/steve-care-software/steve/domain/connections"
	"github.com/steve-care-software/steve/domain/connections/links"
	"github.com/steve-care-software/steve/domain/points"
)

type inMemoryBuilder struct {
	connectionsBuilder connections.Builder
	connectionBuilder  connections.ConnectionBuilder
	linkBuilder        links.LinkBuilder
	connections        connections.Connections
	points             points.Points
}

func createInMemoryBuilder(
	connectionsBuilder connections.Builder,
	connectionBuilder connections.ConnectionBuilder,
	linkBuilder links.LinkBuilder,
) InMemoryBuilder {
	out := inMemoryBuilder{
		connectionsBuilder: connectionsBuilder,
		connectionBuilder:  connectionBuilder,
		linkBuilder:        linkBuilder,
		connections:        nil,
		points:             nil,
	}

	return &out
}

// Create initializes the builder
func (app *inMemoryBuilder) Create() InMemoryBuilder {
	return createInMemoryBuilder(
		app.connectionsBuilder,
		app.connectionBuilder,
		app.linkBuilder,
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

		// if there is a reverse, add the to as the from:
		link := oneConnection.Link()
		if !link.HasReverse() {
			continue
		}

		name := link.Reverse()
		weight := link.Weight()
		newLink, err := app.linkBuilder.Create().WithName(name).WithWeight(weight).Now()
		if err != nil {
			return nil, err
		}

		from := oneConnection.To()
		to := oneConnection.From()
		newConnection, err := app.connectionBuilder.Create().From(from).To(to).WithLink(newLink).Now()
		if err != nil {
			return nil, err
		}

		keyname = newConnection.From().String()
		if _, ok := mpFromList[keyname]; !ok {
			mpFromList[keyname] = []connections.Connection{}
		}

		mpFromList[keyname] = append(mpFromList[keyname], newConnection)
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
