package updates

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens/pointers"
)

type update struct {
	hash   hash.Hash
	origin pointers.Pointer
	target pointers.Pointer
}

func createUpdate(
	hash hash.Hash,
	origin pointers.Pointer,
	target pointers.Pointer,
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
func (obj *update) Target() pointers.Pointer {
	return obj.target
}
