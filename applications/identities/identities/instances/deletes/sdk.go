package deletes

import (
	"github.com/steve-care-software/steve/domain/accounts/identities"
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/identities/instances/successes/deletes"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/identities/instances/contents/deletes"
)

// Application represents the delete application
type Application interface {
	Execute(delete inputs.Delete, current identities.Identity) (executions.Delete, error)
}
