package authenticates

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/identities/authenticates"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/identities/authenticates"
	"github.com/steve-care-software/steve/domain/stacks"
)

// Application represents the application
type Application interface {
	Execute(identity inputs.Authenticate, stack stacks.Stack) (executions.Authenticate, error)
}
