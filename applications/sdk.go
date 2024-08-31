package applications

import (
	"github.com/steve-care-software/steve/domain/queries"
	"github.com/steve-care-software/steve/domain/routes"
)

// Application represents the application
type Application interface {
	Routes(queries queries.Queries) (routes.Routes, error)
	Route(query queries.Query) (routes.Route, error)
}
