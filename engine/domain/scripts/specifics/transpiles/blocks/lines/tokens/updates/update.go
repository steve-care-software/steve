package updates

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles/blocks/lines/tokens/pointers"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles/blocks/lines/tokens/updates/targets"
	"github.com/steve-care-software/steve/hash"
)

type update struct {
	hash   hash.Hash
	origin pointers.Pointer
	target targets.Target
}

func createUpdate(
	hash hash.Hash,
	origin pointers.Pointer,
	target targets.Target,
) Update {
	out := update{
		hash:   hash,
		origin: origin,
		target: target,
	}

	return &out
}

// Hash returns the hash
func (obj *update) Hash() hash.Hash {
	return obj.hash
}

// Origin returns the origin
func (obj *update) Origin() pointers.Pointer {
	return obj.origin
}

// Target returns the target
func (obj *update) Target() targets.Target {
	return obj.target
}
