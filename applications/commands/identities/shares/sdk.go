package shares

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/identities/shares"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/identities/shares"
)

// Application represents the application
type Application interface {
	Execute(shares inputs.Shares, frame frames.Frame) (executions.Shares, error)
}
