package pointers

import (
	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens/pointers/elements"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
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
	Element() elements.Element
	Index() uint
}
