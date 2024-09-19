package updates

import (
	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens/pointers"
)

type update struct {
	origin pointers.Pointer
	target pointers.Pointer
}

func createUpdate(
	origin pointers.Pointer,
	target pointers.Pointer,
) Update {
	out := update{
		origin: origin,
		target: target,
	}

	return &out
}

// Origin returns the origin
func (obj *update) Origin() pointers.Pointer {
	return obj.origin
}

// Target returns the target
func (obj *update) Target() pointers.Pointer {
	return obj.target
}
