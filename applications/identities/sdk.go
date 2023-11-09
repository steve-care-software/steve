package identities

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities"
	"github.com/steve-care-software/steve/domain/stacks"
)

// Application represents the administrator's application
type Application interface {
	Execute(administrator inputs.Identity, stack stacks.Stack) (executions.Identity, error)
}
