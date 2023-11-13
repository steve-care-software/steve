package visitors

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/visitors"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/visitors"
)

// Application represents the visitor application
type Application interface {
	Execute(visitor inputs.Visitor) (executions.Visitor, error)
}
