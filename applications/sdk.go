package applications

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/signers/signatures"
	"github.com/steve-care-software/steve/domain/stencils"
	"github.com/steve-care-software/steve/domain/stencils/messages"
	"github.com/steve-care-software/steve/domain/stencils/results/executions"
)

// Application represents the stencil application
type Application interface {
	Authorize(message messages.Message, username string, password []byte) (executions.Execution, error)
	Authenticate(message messages.Message, signature signatures.Signature) (executions.Execution, error)
	Visit(message messages.Message, stencil stencils.Stencil) (executions.Execution, error)
}
