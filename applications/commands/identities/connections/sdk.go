package connections

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/identities/connections"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/identities/connections"
)

// Application represents the application
type Application interface {
	Execute(connection inputs.Connections, frame frames.Frame) (executions.Connections, error)
}
