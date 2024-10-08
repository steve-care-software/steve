package executions

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/pipelines/executions/suites"
	"github.com/steve-care-software/steve/hash"
)

// Builder represents the execution builder
type Builder interface {
	Create() Builder
	WithRecipe(recipe string) Builder
	WithProgram(program string) Builder
	WithNext(next Execution) Builder
	WithSuites(suites suites.Suites) Builder
	Now() (Execution, error)
}

// Execution represents a pipeline execution
type Execution interface {
	Hash() hash.Hash
	Recipe() string
	HasProgram() bool
	Program() string
	HasNext() bool
	Next() Execution
	HasSuites() bool
	Suites() suites.Suites
}
