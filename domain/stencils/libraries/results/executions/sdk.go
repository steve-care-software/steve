package executions

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/results/executions/actions"
)

// Executions represents executions
type Executions interface {
	List() []Execution
}

// Execution represents an execution
type Execution interface {
	Name() string
	Hash() hash.Hash
	Action() actions.Action
}
