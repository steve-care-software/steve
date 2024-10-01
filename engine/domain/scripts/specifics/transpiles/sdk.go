package transpiles

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles/blocks"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the transpile builder
type Builder interface {
	Create() Builder
	WithBlocks(blocks blocks.Blocks) Builder
	WithRoot(root string) Builder
	WithOrigin(origin string) Builder
	WithTarget(target string) Builder
	Now() (Transpile, error)
}

// Transpile represents a transpile
type Transpile interface {
	Hash() hash.Hash
	Blocks() blocks.Blocks
	Root() string
	Origin() string
	Target() string
}
