package bytevalues

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/bytevalues"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/frames"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/symbols/layers"
)

// Application represents the application
type Application interface {
	Values(values layers.ByteValues, frame frames.Frame) (executions.ByteValues, error)
	Value(value layers.ByteValue, frame frames.Frame) (executions.ByteValue, error)
}
