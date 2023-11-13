package fetches

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/shares/dashboards/successes/fetches"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/shares/dashboards/contents/fetches"
	"github.com/steve-care-software/steve/domain/dashboards"
)

// Application represents the dashboard's fetch application
type Application interface {
	Execute(fetch inputs.Fetch, current dashboards.Dashboard) (executions.Fetch, error)
}
