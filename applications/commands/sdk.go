package commands

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/frames"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs"
)

// Application represents the command application
type Application interface {
	Execute(input inputs.Input, frame frames.Frame) (executions.Execution, error)
}
