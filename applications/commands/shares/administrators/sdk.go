package administrators

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/shares/administrators"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/shares/administrators"
)

// Application represents the application
type Application interface {
	Execute(administrator inputs.Administrator) (executions.Administrator, error)
}
