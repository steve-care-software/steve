package links

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links/executions"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links/origins"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links/preparations"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links/suites"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a link builder
type Builder interface {
	Create() Builder
	WithOrigins(origins origins.Origins) Builder
	WithExecution(execution executions.Execution) Builder
	WithPreparations(preparations preparations.Preparations) Builder
	WithSuites(suites suites.Suites) Builder
	Now() (Link, error)
}

// Link represents a link
type Link interface {
	Hash() hash.Hash
	Origins() origins.Origins
	Execution() executions.Execution
	Preparations() preparations.Preparations
	HasSuites() bool
	Suites() suites.Suites
}
