package passwords

import (
	"github.com/steve-care-software/steve/domain/accounts/identities"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/identities/identities/instances/successes/updates/passwords"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/identities/identities/instances/contents/updates/passwords"
)

// Application represents the application
type Application interface {
	Execute(password inputs.Password, current identities.Identity) (executions.Password, error)
}
