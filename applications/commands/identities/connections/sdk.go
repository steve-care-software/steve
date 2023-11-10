package connections

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/identities/connections"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/identities/connections"
)

// Application represents the application
type Application interface {
	Execute(connection inputs.Connections, frame frames.Frame) (executions.Connections, error)
}
