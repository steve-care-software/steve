package applications

import (
	"github.com/google/uuid"
	applications_connections "github.com/steve-care-software/steve/applications/connections"
	"github.com/steve-care-software/steve/domain/connections"
	"github.com/steve-care-software/steve/domain/routes"
)

// NewApplication creates a new application
func NewApplication(
	connApp applications_connections.Application,
) Application {
	routesBuilder := routes.NewBuilder()
	routeBuilder := routes.NewRouteBuilder()
	connectionsBuilder := connections.NewBuilder()
	return createApplication(
		connApp,
		routesBuilder,
		routeBuilder,
		connectionsBuilder,
	)
}

// Application represents the application
type Application interface {
	// Route returns the possible routes between 2 points
	Route(from uuid.UUID, to uuid.UUID) (routes.Route, error)
}
