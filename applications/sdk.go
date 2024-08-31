package applications

import (
	applications_connections "github.com/steve-care-software/steve/applications/connections"
	"github.com/steve-care-software/steve/domain/connections"
	"github.com/steve-care-software/steve/domain/queries"
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
	Routes(queries queries.Queries) (routes.Routes, error)
	Route(query queries.Query) (routes.Route, error)
}
