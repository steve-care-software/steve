package applications

import (
	"github.com/google/uuid"
	applications_connections "github.com/steve-care-software/steve/applications/connections"
	"github.com/steve-care-software/steve/domain/connections"
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

// Route returns the possible routes between 2 points
func (app *application) Route(from uuid.UUID, to uuid.UUID) (routes.Route, error) {
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

// LinkIntersect intersects the link names and discover their edges
func (app *application) LinkIntersect(names []string) ([]uuid.UUID, error) {
	values := map[string][]uuid.UUID{}
	for _, oneName := range names {
		retIDs, err := app.connApp.ListFromLinkName(oneName)
		if err != nil {
			return nil, err
		}

		values[oneName] = retIDs
	}

	return app.intersect(values)
}

func (app *application) intersect(values map[string][]uuid.UUID) ([]uuid.UUID, error) {
	output := []uuid.UUID{}
	joined := app.join(values)
	expectedLength := uint(len(values))
	for idAsStr, amount := range joined {
		if amount != expectedLength {
			continue
		}

		id, err := uuid.Parse(idAsStr)
		if err != nil {
			return nil, err
		}

		output = append(output, id)
	}

	return output, nil
}

func (app *application) join(values map[string][]uuid.UUID) map[string]uint {
	output := map[string]uint{}
	for _, oneList := range values {
		for _, oneID := range oneList {
			keyname := oneID.String()
			if _, ok := output[keyname]; ok {
				output[keyname]++
				continue
			}

			output[keyname] = 1
		}
	}

	return output
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
