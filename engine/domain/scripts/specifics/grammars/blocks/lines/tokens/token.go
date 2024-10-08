package tokens

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/steve/hash"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens/cardinalities"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens/reverses"
)

type token struct {
	hash        hash.Hash
	element     elements.Element
	cardinality cardinalities.Cardinality
	reverse     reverses.Reverse
}

func createToken(
	hash hash.Hash,
	element elements.Element,
	cardinality cardinalities.Cardinality,
) Token {
	return createTokenInternally(hash, element, cardinality, nil)
}

func createTokenWithReverse(
	hash hash.Hash,
	element elements.Element,
	cardinality cardinalities.Cardinality,
	reverse reverses.Reverse,
) Token {
	return createTokenInternally(hash, element, cardinality, reverse)
}

func createTokenInternally(
	hash hash.Hash,
	element elements.Element,
	cardinality cardinalities.Cardinality,
	reverse reverses.Reverse,
) Token {
	out := token{
		hash:        hash,
		element:     element,
		cardinality: cardinality,
		reverse:     reverse,
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

// Cardinality returns the cardinality
func (obj *token) Cardinality() cardinalities.Cardinality {
	return obj.cardinality
}

// HasReverse returns true if there is a reverse, false ortherwise
func (obj *token) HasReverse() bool {
	return obj.reverse != nil
}

// Reverse returns the reverse, if any
func (obj *token) Reverse() reverses.Reverse {
	return obj.reverse
}
