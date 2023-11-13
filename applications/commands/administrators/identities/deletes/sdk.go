package deletes

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators/identities"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/identities/successes/deletes"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/administrators/identities/contents/deletes"
)

// Application represents the identities application
type Application interface {
	Execute(instance inputs.Delete, current identities.Identities) (executions.Delete, error)
}
