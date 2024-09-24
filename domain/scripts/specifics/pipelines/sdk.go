package pipelines

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/domain/scripts/specifics/pipelines/executions"
)

// Builder represents the pipeline builder
type Builder interface {
	Create() Builder
	WithHead(head heads.Head) Builder
	WithExecution(execution executions.Execution) Builder
	Now() (Pipeline, error)
}

// Pipeline represents the pipeline
type Pipeline interface {
	Hash() hash.Hash
	Head() heads.Head
	Execution() executions.Execution
}
