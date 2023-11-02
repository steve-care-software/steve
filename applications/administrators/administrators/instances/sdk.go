package instances

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators/administrators/instances"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators/instances"
	"github.com/steve-care-software/steve/domain/stacks"
)

// Application represents the instance application
type Application interface {
	Execute(instance inputs.Instance, stack stacks.Stack) (executions.Instance, error)
}
