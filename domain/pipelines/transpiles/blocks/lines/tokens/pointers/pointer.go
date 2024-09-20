package pointers

import (
	"github.com/steve-care-software/steve/domain/pipelines/transpiles/blocks/lines/tokens/pointers/elements"
)

type pointer struct {
	element elements.Element
	index   uint
}

func createPointer(
	element elements.Element,
	index uint,
) Pointer {
	out := pointer{
		element: element,
		index:   index,
	}

	return &out
}

// Element returns the element
func (obj *pointer) Element() elements.Element {
	return obj.element
}

// Index returns the index
func (obj *pointer) Index() uint {
	return obj.index
}
