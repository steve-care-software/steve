package administrators

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/visitors/administrators"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/visitors/administrators"
)

// Application represents the visitor application
type Application interface {
	Execute(administrator inputs.Administrator) (executions.Administrator, error)
}
