package updates

import (
	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens/pointers"
	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens/updates/targets"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an update builder
type Builder interface {
	Create() Builder
	WithOrigin(origin pointers.Pointer) Builder
	WithTarget(target targets.Target) Builder
	Now() (Update, error)
}

// Update represents an update
type Update interface {
	Origin() pointers.Pointer
	Target() targets.Target
}
