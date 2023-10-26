package connections

import (
	"github.com/steve-care-software/steve/applications/connections/authenticates"
	"github.com/steve-care-software/steve/domain/credentials/connections"
)

// Application represents a connection application
type Application interface {
	Initialize() ([]byte, error)
	Authenticate(connection connections.Connection) (authenticates.Application, error)
}
