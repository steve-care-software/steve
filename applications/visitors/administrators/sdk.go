package administrators

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators"
)

// Application represents the visitor application
type Application interface {
	Execute(administrator inputs.Administrator) (executions.Administrator, error)
}
