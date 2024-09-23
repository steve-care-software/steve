package updates

import (
	"github.com/steve-care-software/steve/domain/scripts/specifics/transpiles/blocks/lines/tokens/pointers"
	"github.com/steve-care-software/steve/domain/scripts/specifics/transpiles/blocks/lines/tokens/updates/targets"
)

type update struct {
	origin pointers.Pointer
	target targets.Target
}

func createUpdate(
	origin pointers.Pointer,
	target targets.Target,
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
func (obj *update) Target() targets.Target {
	return obj.target
}
