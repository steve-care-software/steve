package applications

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/queries"
	"github.com/steve-care-software/steve/databases/graphs/domain/responses"
)

// Application represents the graphdb application
type Application interface {
	Execute(qury queries.Query) (responses.Response, error)
}
