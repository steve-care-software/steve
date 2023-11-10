package creates

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/shares/administrators/creates"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/shares/administrators/creates"
)

// Application represents the application
type Application interface {
	Execute(administrator inputs.Create) (executions.Create, error)
}
