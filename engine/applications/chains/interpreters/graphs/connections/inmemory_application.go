package connections

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/steve-care-software/steve/engine/domain/graphs/connections"
)

type application struct {
	fromConnections         map[string]connections.Connections
	mpPointsToLinkNameEdges map[string][]uuid.UUID
}

func createApplication(
	fromConnections map[string]connections.Connections,
	mpPointsToLinkNameEdges map[string][]uuid.UUID,
) Application {
	out := application{
		fromConnections:         fromConnections,
		mpPointsToLinkNameEdges: mpPointsToLinkNameEdges,
	}

	return &out
}

// ListFrom lists the connections by its from UUID
func (app *application) ListFrom(from uuid.UUID) (connections.Connections, error) {
	keyname := from.String()
	if list, ok := app.fromConnections[keyname]; ok {
		return list, nil
	}

	str := fmt.Sprintf("there is no connections that starts at the provided identifier: %s", from.String())
	return nil, errors.New(str)
}

// ListFromLinkName returns the list of ids that touch the link
func (app *application) ListFromLinkName(name string) ([]uuid.UUID, error) {
	if list, ok := app.mpPointsToLinkNameEdges[name]; ok {
		return list, nil
	}

	str := fmt.Sprintf("the link (%s) does not have any connected edge", name)
	return nil, errors.New(str)
}
