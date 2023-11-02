package fetches

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators/instances/fetches"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators/instances/contents/fetches"
)

// Application represents the fetch application
type Application interface {
	Execute(fetch inputs.Fetch) (executions.Fetch, error)
}
