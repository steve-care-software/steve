package instances

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators/administrators/instances"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators/instances"
)

// Application represents the instance application
type Application interface {
	Execute(instance inputs.Instance) (executions.Instance, error)
}
