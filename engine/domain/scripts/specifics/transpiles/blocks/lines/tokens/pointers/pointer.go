package pointers

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles/blocks/lines/tokens/pointers/elements"
)

type pointer struct {
	hash    hash.Hash
	element elements.Element
	index   uint
}

func createPointer(
	hash hash.Hash,
	element elements.Element,
	index uint,
) Pointer {
	out := pointer{
		hash:    hash,
		element: element,
		index:   index,
	}

	return &out
}

// Hash returns the hash
func (obj *pointer) Hash() hash.Hash {
	return obj.hash
}

// Element returns the element
func (obj *pointer) Element() elements.Element {
	return obj.element
}

// Index returns the index
func (obj *pointer) Index() uint {
	return obj.index
}
