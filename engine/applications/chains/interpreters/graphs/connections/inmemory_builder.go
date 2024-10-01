package connections

import (
	"errors"

	"github.com/google/uuid"
	"github.com/steve-care-software/steve/engine/domain/graphs/connections"
	"github.com/steve-care-software/steve/engine/domain/graphs/connections/links"
	"github.com/steve-care-software/steve/engine/domain/graphs/points"
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

	mpPointsToLinkNameEdges := map[string][]uuid.UUID{}
	mpFromList := map[string][]connections.Connection{}
	list := app.connections.List()
	for _, oneConnection := range list {
		keyname := oneConnection.From().String()
		if _, ok := mpFromList[keyname]; !ok {
			mpFromList[keyname] = []connections.Connection{}
		}

		mpFromList[keyname] = append(mpFromList[keyname], oneConnection)

		// link names:
		link := oneConnection.Link()
		linkName := link.Name()
		if _, ok := mpPointsToLinkNameEdges[linkName]; !ok {
			mpPointsToLinkNameEdges[linkName] = []uuid.UUID{}
		}

		mpPointsToLinkNameEdges[linkName] = append(mpPointsToLinkNameEdges[linkName], oneConnection.From())

		// if there is a reverse, add the to as the from:
		if !link.HasReverse() {
			continue
		}

		name := link.Reverse()
		newLink, err := app.linkBuilder.Create().WithName(name).Now()
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

		// reverse names:
		reverseLinkName := newLink.Name()
		if _, ok := mpPointsToLinkNameEdges[reverseLinkName]; !ok {
			mpPointsToLinkNameEdges[reverseLinkName] = []uuid.UUID{}
		}

		mpPointsToLinkNameEdges[reverseLinkName] = append(mpPointsToLinkNameEdges[reverseLinkName], oneConnection.From())
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
		mpPointsToLinkNameEdges,
	), nil
}
