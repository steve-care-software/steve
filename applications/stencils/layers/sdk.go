package layers

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/executions"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/frames"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/symbols/layers"
)

// Application represents the application
type Application interface {
	Execute(layer layers.Layer, frame frames.Frame) (executions.Execution, error)
}
