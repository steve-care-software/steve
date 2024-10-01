package executions

import (
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/executions/parameters"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an execution builder
type Builder interface {
	Create() Builder
	WithParameters(parameters parameters.Parameters) Builder
	WithFuncName(fnFlag string) Builder
	Now() (Execution, error)
}

// Execution represents an execution
type Execution interface {
	FuncName() string
	HasParameters() bool
	Parameters() parameters.Parameters
}