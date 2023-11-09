package administrators

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/shares/administrators"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/shares/administrators"
)

// Application represents the application
type Application interface {
	Execute(administrator inputs.Administrator) (executions.Administrator, error)
}
