package instances

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/identities/instances"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/identities/instances"
	"github.com/steve-care-software/steve/domain/stacks"
)

// Application represents the application
type Application interface {
	Execute(identity inputs.Instance, stack stacks.Stack) (executions.Instance, error)
}
