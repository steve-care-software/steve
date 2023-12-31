package instances

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/administrators/instances"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/administrators/administrators/instances"
)

// Application represents the instance application
type Application interface {
	Execute(instance inputs.Instance, frame frames.Frame) (executions.Instance, error)
}
