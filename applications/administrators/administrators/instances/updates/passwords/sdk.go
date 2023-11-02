package passwords

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators/instances/updates/passwords"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators/instances/contents/updates/passwords"
)

// Application represents the fetch application
type Application interface {
	Execute(password inputs.Password, instance administrators.Administrator) (executions.Password, error)
}
