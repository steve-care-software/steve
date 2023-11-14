package instructions

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/executions/instructions"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/frames"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/symbols/layers"
)

// Application represents the application
type Application interface {
	Instructions(linstructions layers.Instructions, frame frames.Frame) (executions.Instructions, error)
	Instruction(instruction layers.Instruction, frame frames.Frame) (executions.Instruction, error)
}
