package roots

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/pipelines/executions/suites"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the root builder
type Builder interface {
	Create() Builder
	WithHead(head heads.Head) Builder
	WithPipelines(pipelines []string) Builder
	WithSuites(suites suites.Suites) Builder
	Now() (Root, error)
}

// Root represents a root
type Root interface {
	Hash() hash.Hash
	Head() heads.Head
	Pipelines() []string
	HasSuites() bool
	Suites() suites.Suites
}
