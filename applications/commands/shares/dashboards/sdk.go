package dashboards

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/shares/dashboards"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/shares/dashboards"
)

// Application represents the dashboard's application
type Application interface {
	Execute(dashboard inputs.Dashboard, frame frames.Frame) (executions.Dashboard, error)
}
