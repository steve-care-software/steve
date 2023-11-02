package administrators

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators/administrators"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators"
	"github.com/steve-care-software/steve/domain/stacks"
)

// Application represents the instance application
type Application interface {
	Execute(instance inputs.Administrator, stack stacks.Stack) (executions.Administrator, error)
}
