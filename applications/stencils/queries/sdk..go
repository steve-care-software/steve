package queries

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/executions/queries"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/frames"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/symbols/layers"
)

// Application represents the application
type Application interface {
	Execute(query layers.Query, frame frames.Frame) (executions.Query, error)
}
