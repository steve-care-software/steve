package chains

import (
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens/elements"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewTokenBuilder creates a new token builder
func NewTokenBuilder() TokenBuilder {
	return createTokenBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// Builder represents a chain builder
type Builder interface {
	Create() Builder
	WithElement(element elements.Element) Builder
	WithToken(token Token) Builder
	Now() (Chain, error)
}

// Chain represents a chain
type Chain interface {
	Element() elements.Element
	HasToken() bool
	Token() Token
}

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithIndex(index uint) TokenBuilder
	WithElement(element Element) TokenBuilder
	Now() (Token, error)
}

// Token represents the token
type Token interface {
	Index() uint
	HasElement() bool
	Element() Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithIndex(index uint) ElementBuilder
	WithChain(chain Chain) ElementBuilder
	Now() (Element, error)
}

// Element represents the element
type Element interface {
	Index() uint
	HasChain() bool
	Chain() Chain
}
