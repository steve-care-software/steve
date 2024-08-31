package applications

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	applications_connections "github.com/steve-care-software/steve/applications/connections"
	"github.com/steve-care-software/steve/domain/connections"
	"github.com/steve-care-software/steve/domain/paths"
	"github.com/steve-care-software/steve/domain/queries"
)

type application struct {
	connApp            applications_connections.Application
	pathsBuilder       paths.Builder
	connectionsBuilder connections.Builder
}

func createApplication(
	connApp applications_connections.Application,
	pathsBuilder paths.Builder,
	connectionsBuilder connections.Builder,
) Application {
	out := application{
		connApp:            connApp,
		pathsBuilder:       pathsBuilder,
		connectionsBuilder: connectionsBuilder,
	}

	return &out
}

// Discover discovers the path of a query
func (app *application) Discover(queries queries.Queries) (paths.Paths, error) {
	list := queries.List()
	pathList := []connections.Connections{}
	for _, oneQuery := range list {
		retConnections, err := app.discoverQuery(oneQuery)
		if err != nil {
			return nil, err
		}

		pathList = append(pathList, retConnections)
	}

	return app.pathsBuilder.Create().
		WithList(pathList).
		Now()
}

func (app *application) discoverQuery(query queries.Query) (connections.Connections, error) {
	from := query.From()
	to := query.To()
	return app.followUntilReached(
		from,
		to,
		[]connections.Connection{},
	)
}

func (app *application) followUntilReached(
	start uuid.UUID,
	destination uuid.UUID,
	connectionsList []connections.Connection,
) (connections.Connections, error) {
	listTo, err := app.connApp.ListFrom(start)
	if err != nil {
		str := fmt.Sprintf("there is no link between the requested points (start: %s, to: %s)", start.String(), destination.String())
		return nil, errors.New(str)
	}

	listToConnectionsList := listTo.List()
	for _, oneConnection := range listToConnectionsList {
		if oneConnection.To().String() == destination.String() {
			// reached destination
			merged := append(connectionsList, oneConnection)
			return app.connectionsBuilder.Create().
				WithList(merged).
				Now()
		}

		newFrom := oneConnection.To()
		merged := append(connectionsList, oneConnection)
		retConnections, err := app.followUntilReached(
			newFrom,
			destination,
			merged,
		)

		if err != nil {
			continue
		}

		return retConnections, nil
	}

	str := fmt.Sprintf("there is no path between the requested points (start: %s, to: %s)", start.String(), destination.String())
	return nil, errors.New(str)
}
