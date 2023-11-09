package shares

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/shares"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/shares"
	"github.com/steve-care-software/steve/domain/stacks"
)

// Application represents the application
type Application interface {
	Execute(shares inputs.Shares, stack stacks.Stack) (executions.Shares, error)
}
