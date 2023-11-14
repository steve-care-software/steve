package dashboards

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/dashboards/stencils"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the dashboard hash
type Builder interface {
	Create() Builder
	WithRoot(root hash.Hash) Builder
	WithVisitor(visitor hash.Hash) Builder
	WithLibrary(library stencils.Stencils) Builder
	Now() (Dashboard, error)
}

// Dashboard represents a dashboard
type Dashboard interface {
	Root() stencils.Stencil
	Visitor() stencils.Stencil
	Library() stencils.Stencils
}
