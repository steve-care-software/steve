package authenticates

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators/authenticates"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators/authenticates"
)

// Application represents the administrator's authenticate pplication
type Application interface {
	Execute(administrator inputs.Authenticate) (executions.Authenticate, error)
}
