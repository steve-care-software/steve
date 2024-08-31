package applications

import (
	"github.com/google/uuid"
	applications_connections "github.com/steve-care-software/steve/applications/connections"
	"github.com/steve-care-software/steve/domain/connections"
	"github.com/steve-care-software/steve/domain/queries"
	"github.com/steve-care-software/steve/domain/routes"
)

type application struct {
	connApp            applications_connections.Application
	routesBuilder      routes.Builder
	routeBuilder       routes.RouteBuilder
	connectionsBuilder connections.Builder
}

func createApplication(
	connApp applications_connections.Application,
	routesBuilder routes.Builder,
	routeBuilder routes.RouteBuilder,
	connectionsBuilder connections.Builder,
) Application {
	out := application{
		connApp:            connApp,
		routesBuilder:      routesBuilder,
		routeBuilder:       routeBuilder,
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
	retConnectionsList, err := app.followUntilReached(
		from,
		to,
		[]connections.Connection{},
	)

	if err != nil {
		return nil, err
	}

	return app.routeBuilder.Create().
		WithPossibilities(retConnectionsList).
		Now()
}

func (app *application) followUntilReached(
	start uuid.UUID,
	destination uuid.UUID,
	connectionsList []connections.Connection,
) ([]connections.Connections, error) {
	listTo, err := app.connApp.ListFrom(start)
	if err != nil {
		return nil, err
	}

	retOutputList := []connections.Connections{}
	listToConnectionsList := listTo.List()
	for _, oneConnection := range listToConnectionsList {
		// make sure there is no circular link:
		skip := false
		currentToStr := oneConnection.To().String()
		for _, onePreviousConn := range connectionsList {
			prevFromStr := onePreviousConn.From().String()
			if prevFromStr == currentToStr {
				skip = true
				break
			}
		}

		if skip {
			continue
		}

		merged := append(connectionsList, oneConnection)
		if oneConnection.To().String() == destination.String() {
			// reached destination
			retConnections, err := app.connectionsBuilder.Create().
				WithList(merged).
				Now()

			if err != nil {
				return nil, err
			}

			retOutputList = append(retOutputList, retConnections)
			continue
		}

		newFrom := oneConnection.To()
		retConnectionsList, err := app.followUntilReached(
			newFrom,
			destination,
			merged,
		)

		if err != nil {
			continue
		}

		retOutputList = append(retOutputList, retConnectionsList...)
	}

	return retOutputList, nil
}
