package visitors

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/visitors"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/visitors"
)

// Application represents the visitor application
type Application interface {
	Execute(visitor inputs.Visitor) (executions.Visitor, error)
}
