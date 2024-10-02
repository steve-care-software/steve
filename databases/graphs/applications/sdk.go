package applications

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/links"
	"github.com/steve-care-software/steve/databases/graphs/domain/routes"
)

// Application represents the graphdb application
type Application interface {
	Route(link links.Link) (routes.Routes, error)
}
