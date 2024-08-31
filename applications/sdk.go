package applications

import (
	"github.com/steve-care-software/steve/domain/paths"
	"github.com/steve-care-software/steve/domain/queries"
)

// Application represents the application
type Application interface {
	Discover(queries queries.Queries) (paths.Paths, error)
}
