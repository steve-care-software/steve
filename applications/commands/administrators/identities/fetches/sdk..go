package fetches

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators/identities"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/identities/successes/fetches"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/administrators/identities/contents/fetches"
)

// Application represents the fetch application
type Application interface {
	Execute(instance inputs.Fetch, current identities.Identities) (executions.Fetch, error)
}
