package applications

import (
	"github.com/steve-care-software/steve/graphs/domain/responses"
	"github.com/steve-care-software/steve/graphs/domain/scripts"
	"github.com/steve-care-software/steve/hash"
	application_lists "github.com/steve-care-software/steve/lists/applications"
	application_resources "github.com/steve-care-software/steve/resources/applications"
)

// NewBuilder creates a new builder instance
func NewBuilder(
	storeListApp application_lists.Application,
	resourceApp application_resources.Application,
) Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		storeListApp,
		resourceApp,
		hashAdapter,
	)
}

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithIdentifier(dbIdentifier string) Builder
	Now() (Application, error)
}

// Application represents the graphdb application
type Application interface {
	Execute(script scripts.Script) (responses.Response, error)
}
