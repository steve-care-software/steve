package applications

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	"github.com/steve-care-software/steve/domain/accounts/identities"
	"github.com/steve-care-software/steve/domain/accounts/identities/profiles"
	"github.com/steve-care-software/steve/domain/stencils"
	"github.com/steve-care-software/steve/domain/stencils/messages"
	"github.com/steve-care-software/steve/domain/stencils/results/executions"
)

// Application represents the stencil application
type Application interface {
	Administrator(message messages.Message, administrator administrators.Administrator) (executions.Execution, error)
	Identity(message messages.Message, stencil stencils.Stencil, identity identities.Identity) (executions.Execution, error)
	Connection(message messages.Message, stencil stencils.Stencil, owner identities.Identity, profile profiles.Connection) (executions.Execution, error)
	Visitor(message messages.Message, stencil stencils.Stencil) (executions.Execution, error)
}
