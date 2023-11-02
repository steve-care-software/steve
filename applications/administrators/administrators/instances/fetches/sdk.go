package fetches

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators/instances/fetches"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators/instances/contents/fetches"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithInstance(instance administrators.Administrator) Builder
	Now() (Application, error)
}

// Application represents the fetch application
type Application interface {
	Execute(fetch inputs.Fetch) (executions.Fetch, error)
}
