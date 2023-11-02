package identities

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators/identities"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators/identities"
	"github.com/steve-care-software/steve/domain/stacks"
)

// Application represents the identities application
type Application interface {
	Execute(instance inputs.Identities, stack stacks.Stack) (executions.Identities, error)
}
