package results

import "github.com/steve-care-software/steve/domain/stencils/libraries/results/executions"

// Result represents a library save result
type Result interface {
	Executions() executions.Executions
}
