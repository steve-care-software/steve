package assignables

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/executions/instructions/assignments/assignables"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/frames"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/symbols/layers"
)

const trueByte = 1
const falseByte = 0

// Application represents the application
type Application interface {
	Execute(assignable layers.Assignable, frame frames.Frame) (executions.Assignable, error)
}
