package signers

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/signers"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/signers"
	"github.com/steve-care-software/steve/domain/stacks"
)

// Application represents the application
type Application interface {
	Execute(signer inputs.Signer, stack stacks.Stack) (executions.Signer, error)
}
