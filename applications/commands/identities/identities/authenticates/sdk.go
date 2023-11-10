package authenticates

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/identities/identities/authenticates"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/identities/identities/authenticates"
)

// Application represents the application
type Application interface {
	Execute(authenticate inputs.Authenticate) (executions.Authenticate, error)
}
