package encryptors

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/encryptors"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/encryptors"
	"github.com/steve-care-software/steve/domain/stacks"
)

// Application represents the application
type Application interface {
	Execute(encryptor inputs.Encryptor, stack stacks.Stack) (executions.Encryptor, error)
}
