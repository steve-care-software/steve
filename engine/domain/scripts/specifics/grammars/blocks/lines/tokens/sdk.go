package tokens

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens/cardinalities"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens/reverses"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewTokenBuilder creates a new token builder
func NewTokenBuilder() TokenBuilder {
	hashAdpater := hash.NewAdapter()
	return createTokenBuilder(
		hashAdpater,
	)
}

// Builder represents the tokens builder
type Builder interface {
	Create() Builder
	WithList(list []Token) Builder
	Now() (Tokens, error)
}

// Tokens represents tokens
type Tokens interface {
	Hash() hash.Hash
	List() []Token
}

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithElement(element elements.Element) TokenBuilder
	WithCardinality(cardinality cardinalities.Cardinality) TokenBuilder
	WithReverse(reverse reverses.Reverse) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Hash() hash.Hash
	Element() elements.Element
	Cardinality() cardinalities.Cardinality
	HasReverse() bool
	Reverse() reverses.Reverse
}