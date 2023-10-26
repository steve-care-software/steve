package executions

import (
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns"
	"github.com/steve-care-software/steve/domain/stencils/results/executions/executions"
	"github.com/steve-care-software/steve/domain/stencils/results/executions/inits"
)

// Builder represents the execution results builder
type Builder interface {
	Create() Builder
	WithList(list []Execution) Builder
	Now() (Executions, error)
}

// Executions represents executions
type Executions interface {
	List() []Execution
}

// ExecutionBuilder represents a execution result builder
type ExecutionBuilder interface {
	Create() ExecutionBuilder
	WithInit(init inits.Init) ExecutionBuilder
	WithExecutions(executions executions.Executions) ExecutionBuilder
	WithReturn(ret returns.Return) ExecutionBuilder
	Now() (Execution, error)
}

// Execution represents execution result
type Execution interface {
	Init() inits.Init
	Executions() executions.Executions
	Return() returns.Return
	Bytes() []byte
}
