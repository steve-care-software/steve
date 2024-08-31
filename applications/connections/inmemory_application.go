package connections

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/connections"
)

type application struct {
	fromConnections map[string]connections.Connections
}

func createApplication(
	fromConnections map[string]connections.Connections,
) Application {
	out := application{
		fromConnections: fromConnections,
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
