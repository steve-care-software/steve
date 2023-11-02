package deletes

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators/instances/deletes"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators/instances/contents/deletes"
)

// Application represents the delete application
type Application interface {
	Execute(delete inputs.Delete) (executions.Delete, error)
}
