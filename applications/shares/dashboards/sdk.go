package dashboards

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/shares/dashboards"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/shares/dashboards"
	"github.com/steve-care-software/steve/domain/stacks"
)

// Application represents the dashboard's application
type Application interface {
	Execute(dashboard inputs.Dashboard, stack stacks.Stack) (executions.Dashboard, error)
}
