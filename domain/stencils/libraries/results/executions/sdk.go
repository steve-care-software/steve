package executions

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/results/executions/actions"
)

// Builder represents an executions builder
type Builder interface {
	Create() Builder
	WithList(list []Execution) Builder
	Now() (Executions, error)
}

// Executions represents executions
type Executions interface {
	List() []Execution
}

// ExecutionBuilder represents an execution builder
type ExecutionBuilder interface {
	Create() ExecutionBuilder
	WithName(name string) ExecutionBuilder
	WithHash(hash hash.Hash) ExecutionBuilder
	WithAction(action actions.Action) ExecutionBuilder
	Now() (Execution, error)
}

// Execution represents an execution
type Execution interface {
	Name() string
	Hash() hash.Hash
	Action() actions.Action
}
