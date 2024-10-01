package pointers

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles/blocks/lines/tokens/pointers/elements"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the pointer builder
type Builder interface {
	Create() Builder
	WithElement(element elements.Element) Builder
	WithIndex(index uint) Builder
	Now() (Pointer, error)
}

// Pointer represents an element pointer
type Pointer interface {
	Hash() hash.Hash
	Element() elements.Element
	Index() uint
}
