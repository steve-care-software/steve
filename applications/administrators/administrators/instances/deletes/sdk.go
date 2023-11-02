package deletes

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators/administrators/instances/successes/deletes"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators/instances/contents/deletes"
)

// Application represents the delete application
type Application interface {
	Execute(delete inputs.Delete, current administrators.Administrator) (executions.Delete, error)
}
