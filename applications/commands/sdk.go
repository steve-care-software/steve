package commands

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/frames"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs"
)

// Application represents the command application
type Application interface {
	Execute(input inputs.Input, frame frames.Frame) (executions.Execution, error)
}
