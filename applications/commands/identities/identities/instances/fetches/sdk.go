package fetches

import (
	"github.com/steve-care-software/steve/domain/accounts/identities"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/identities/identities/instances/successes/fetches"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/identities/identities/instances/contents/fetches"
)

// Application represents the fetch application
type Application interface {
	Execute(fetch inputs.Fetch, current identities.Identity) (executions.Fetch, error)
}
