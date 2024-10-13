package bridges

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/bridges/links"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the bridge builder
type Builder interface {
	Create() Builder
	WithWeight(weight uint) Builder
	WithOrigin(origin links.Link) Builder
	WithTarget(target links.Link) Builder
	Now() (Bridge, error)
}

// Bridge represents a bridge
type Bridge interface {
	Weight() uint
	Origin() links.Link
	Target() links.Link
}
