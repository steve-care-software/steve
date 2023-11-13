package administrators

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/administrators"
)

// Application represents the administrator's application
type Application interface {
	Execute(administrator inputs.Administrator, frame frames.Frame) (executions.Administrator, error)
}
