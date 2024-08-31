package applications

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	applications_connections "github.com/steve-care-software/steve/applications/connections"
	"github.com/steve-care-software/steve/domain/connections"
	"github.com/steve-care-software/steve/domain/paths"
	"github.com/steve-care-software/steve/domain/queries"
	"github.com/steve-care-software/steve/domain/routes"
)

type application struct {
	connApp            applications_connections.Application
	routesBuilder      routes.Builder
	routeBuilder       routes.RouteBuilder
	pathsBuilder       paths.Builder
	pathBuilder        paths.PathBuilder
	connectionsBuilder connections.Builder
}

func createApplication(
	connApp applications_connections.Application,
	routesBuilder routes.Builder,
	routeBuilder routes.RouteBuilder,
	pathsBuilder paths.Builder,
	pathBuilder paths.PathBuilder,
	connectionsBuilder connections.Builder,
) Application {
	out := application{
		connApp:            connApp,
		routesBuilder:      routesBuilder,
		routeBuilder:       routeBuilder,
		pathsBuilder:       pathsBuilder,
		pathBuilder:        pathBuilder,
		connectionsBuilder: connectionsBuilder,
	}

	return &out
}

// Routes execute queries in order to discover matching routes
func (app *application) Routes(queries queries.Queries) (routes.Routes, error) {
	list := queries.List()
	routeList := []routes.Route{}
	for _, oneQuery := range list {
		retPath, err := app.Route(oneQuery)
		if err != nil {
			return nil, err
		}

		routeList = append(routeList, retPath)
	}

	return app.routesBuilder.Create().
		WithList(routeList).
		Now()
}

// Route execute a query in order to discover a matching route
func (app *application) Route(query queries.Query) (routes.Route, error) {
	from := query.From()
	to := query.To()
	retPath, err := app.followUntilReached(
		from,
		to,
		[]connections.Connection{},
	)

	if err != nil {
		return nil, err
	}

	connectionsList := []connections.Connections{}
	successfuls := retPath.Successfuls()
	for _, oneSuccessful := range successfuls {
		connections, err := app.connectionsBuilder.Create().
			WithList(oneSuccessful).
			Now()

		if err != nil {
			return nil, err
		}

		connectionsList = append(connectionsList, connections)
	}

	return app.routeBuilder.Create().
		WithPossibilities(connectionsList).
		Now()
}

func (app *application) followUntilReached(
	start uuid.UUID,
	destination uuid.UUID,
	connectionsList []connections.Connection,
) (paths.Path, error) {
	listTo, err := app.connApp.ListFrom(start)
	if err != nil {
		str := fmt.Sprintf("there is no link between the requested points (start: %s, to: %s)", start.String(), destination.String())
		return nil, errors.New(str)
	}

	retPathList := []paths.Path{}
	listToConnectionsList := listTo.List()
	for _, oneConnection := range listToConnectionsList {
		if oneConnection.To().String() == destination.String() {
			// reached destination
			paths, err := app.pathsBuilder.Create().WithList(retPathList).Now()
			if err != nil {
				return nil, err
			}

			return app.pathBuilder.Create().
				WithDestination(oneConnection).
				WithPossibilities(paths).
				Now()
		}

		newFrom := oneConnection.To()
		merged := append(connectionsList, oneConnection)
		retPath, err := app.followUntilReached(
			newFrom,
			destination,
			merged,
		)

		if err != nil {
			continue
		}

		retPathList = append(retPathList, retPath)
	}

	if len(retPathList) <= 0 {
		str := fmt.Sprintf("there is no path between the requested points (start: %s, to: %s)", start.String(), destination.String())
		return nil, errors.New(str)
	}

	paths, err := app.pathsBuilder.Create().
		WithList(retPathList).
		Now()

	if err != nil {
		return nil, err
	}

	return app.pathBuilder.Create().
		WithPossibilities(paths).
		Now()
}
