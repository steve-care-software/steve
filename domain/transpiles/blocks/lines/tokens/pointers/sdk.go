package pointers

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens/pointers/elements"
)

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
