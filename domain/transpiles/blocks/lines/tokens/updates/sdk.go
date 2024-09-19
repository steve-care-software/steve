package updates

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens/pointers"
)

// Builder represents an update builder
type Builder interface {
	Create() Builder
	WithOrigin(origin pointers.Pointer) Builder
	WithTarget(target pointers.Pointer) Builder
	Now() (Update, error)
}

// Update represents an update
type Update interface {
	Hash() hash.Hash
	Origin() pointers.Pointer
	Target() pointers.Pointer
}
