package connections

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/connections"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/connections"
	"github.com/steve-care-software/steve/domain/stacks"
)

// Application represents the application
type Application interface {
	Execute(connection inputs.Connections, stack stacks.Stack) (executions.Connections, error)
}
