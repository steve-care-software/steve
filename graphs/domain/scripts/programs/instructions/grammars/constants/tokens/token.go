package tokens

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/grammars/constants/tokens/elements"
	"github.com/steve-care-software/steve/hash"
)

type token struct {
	hash       hash.Hash
	element    elements.Element
	occurences uint
}

func createToken(
	hash hash.Hash,
	element elements.Element,
	occurences uint,
) Token {
	out := token{
		hash:       hash,
		element:    element,
		occurences: occurences,
	}

	return &out
}

// Hash returns the hash
func (obj *token) Hash() hash.Hash {
	return obj.hash
}

// Element returns the element
func (obj *token) Element() elements.Element {
	return obj.element
}

// Occurences returns the occurences
func (obj *token) Occurences() uint {
	return obj.occurences
}
