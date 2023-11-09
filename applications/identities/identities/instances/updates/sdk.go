package updates

import (
	"github.com/steve-care-software/steve/domain/accounts/identities"
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/identities/instances/successes/updates"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/identities/instances/contents/updates"
)

// Application represents the update application
type Application interface {
	Execute(update inputs.Update, current identities.Identity) (executions.Update, error)
}
