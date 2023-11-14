package assignments

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/executions/instructions/assignments"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/frames"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/symbols/layers"
)

// Application represents the application
type Application interface {
	Execute(assignment layers.Assignment, frame frames.Frame) (executions.Assignment, error)
}
