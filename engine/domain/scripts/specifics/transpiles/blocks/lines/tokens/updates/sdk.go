package updates

import (
	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles/blocks/lines/tokens/pointers"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles/blocks/lines/tokens/updates/targets"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
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
	Hash() hash.Hash
	Origin() pointers.Pointer
	Target() targets.Target
}
