package lists

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/identities/lists"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/identities/lists"
	"github.com/steve-care-software/steve/domain/stacks"
)

// Application represents the application
type Application interface {
	Execute(identity inputs.List, stack stacks.Stack) (executions.List, error)
}
