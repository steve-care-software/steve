package encryptors

import (
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/identities/encryptors"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/identities/encryptors"
)

// Application represents the application
type Application interface {
	Execute(encryptor inputs.Encryptor, frame frames.Frame) (executions.Encryptor, error)
}
