package reverses

import (
	"github.com/steve-care-software/steve/hash"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens/elements"
)

type reverse struct {
	hash   hash.Hash
	escape elements.Element
}

func createReverse(
	hash hash.Hash,
) Reverse {
	return createReverseInternally(hash, nil)
}

func createReverseWithEscape(
	hash hash.Hash,
	escape elements.Element,
) Reverse {
	return createReverseInternally(hash, escape)
}

func createReverseInternally(
	hash hash.Hash,
	escape elements.Element,
) Reverse {
	out := reverse{
		hash:   hash,
		escape: escape,
	}

	return &out
}

// Hash returns the hash
func (obj *reverse) Hash() hash.Hash {
	return obj.hash
}

// HasEscape returns true if there is an escape, false otherwise
func (obj *reverse) HasEscape() bool {
	return obj.escape != nil
}

// Escape returns the escape, if any
func (obj *reverse) Escape() elements.Element {
	return obj.escape
}
