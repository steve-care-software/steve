package signers

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/identities/signers"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/identities/signers"
)

// Application represents the application
type Application interface {
	Execute(signer inputs.Signer, frame frames.Frame) (executions.Signer, error)
}
