package instances

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/identities/identities/instances"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/identities/identities/instances"
)

// Application represents the application
type Application interface {
	Execute(identity inputs.Instance, frame frames.Frame) (executions.Instance, error)
}
