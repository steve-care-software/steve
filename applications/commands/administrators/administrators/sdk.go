package administrators

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/administrators"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/administrators/administrators"
)

// Application represents the instance application
type Application interface {
	Execute(instance inputs.Administrator, frame frames.Frame) (executions.Administrator, error)
}
