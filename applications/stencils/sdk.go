package stencils

import (
	"github.com/steve-care-software/steve/domain/stencils"
	"github.com/steve-care-software/steve/domain/stencils/messages"
	"github.com/steve-care-software/steve/domain/stencils/results/executions"
)

// Application represents the stencil application
type Application interface {
	Execute(message messages.Message, stencil stencils.Stencil) (executions.Execution, error)
}
