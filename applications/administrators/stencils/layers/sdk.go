package layers

import (
	"github.com/steve-care-software/steve/domain/stencils/queries"
	result_layers "github.com/steve-care-software/steve/domain/stencils/results/executions"
)

// Application represents a layer application
type Application interface {
	Execute(query queries.Query) (result_layers.Execution, error)
}
