package updates

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/administrators/instances/successes/updates"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/administrators/administrators/instances/contents/updates"
)

// Application represents the update application
type Application interface {
	Execute(update inputs.Update, current administrators.Administrator) (executions.Update, error)
}
