package results

import "github.com/steve-care-software/steve/domain/stencils/libraries/results/executions"

// Builder represents a result builder
type Builder interface {
	Create() Builder
	WithExecutions(executions executions.Executions) Builder
	Now() (Result, error)
}

// Result represents a library save result
type Result interface {
	Executions() executions.Executions
}
