package chains

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens/elements"
)

// Builder represents a chain builder
type Builder interface {
	Create() Builder
	WithElement(element elements.Element) Builder
	WithTokenIndex(tokenIndex uint) Builder
	WithElementIndex(elementIndex uint) Builder
	WithNext(next Chain) Builder
	Now() (Chain, error)
}

// Chain represents a chain
type Chain interface {
	Hash() hash.Hash
	Element() elements.Element
	TokenIndex() uint
	ElementIndex() uint
	HasNext() bool
	Next() Chain
}
