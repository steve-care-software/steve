package fetches

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/administrators/instances/successes/fetches"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/administrators/administrators/instances/contents/fetches"
)

// Application represents the fetch application
type Application interface {
	Execute(fetch inputs.Fetch, current administrators.Administrator) (executions.Fetch, error)
}
