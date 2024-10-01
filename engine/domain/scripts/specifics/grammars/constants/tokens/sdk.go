package tokens

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/grammars/constants/tokens/elements"
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
	hashAdapter := hash.NewAdapter()
	return createTokenBuilder(
		hashAdapter,
	)
}

// Builder represents a tokens builder
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

// TokenBuilder represents the token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithElement(element elements.Element) TokenBuilder
	WithOccurences(occurences uint) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Hash() hash.Hash
	Element() elements.Element
	Occurences() uint
}
