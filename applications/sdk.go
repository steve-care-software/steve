package applications

import (
	"github.com/steve-care-software/steve/domain/commands/executions"
)

// Application represents the stencil application
type Application interface {
	Execute(input []byte) (executions.Execution, error)
}
