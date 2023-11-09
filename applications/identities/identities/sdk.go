package identities

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/identities"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/identities"
	"github.com/steve-care-software/steve/domain/stacks"
)

// Application represents the application
type Application interface {
	Execute(identity inputs.Identity, stack stacks.Stack) (executions.Identity, error)
}
