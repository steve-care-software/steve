package inserts

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators/identities"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/identities/successes/inserts"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/administrators/identities/contents/inserts"
)

// Application represents the fetch application
type Application interface {
	Execute(instance inputs.Insert, current identities.Identities) (executions.Insert, error)
}
