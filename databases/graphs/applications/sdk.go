package applications

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/routes"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/links"
)

// Application represents the graphdb application
type Application interface {
	Route(link links.Link) (routes.Routes, error)
}
