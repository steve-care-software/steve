package selectors

import (
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewSelectorBuilder creates a new selector builder
func NewSelectorBuilder() SelectorBuilder {
	return createSelectorBuilder()
}

// Builder represents the selectors builder
type Builder interface {
	Create() Builder
	WithList(list []Selector) Builder
	Now() (Selectors, error)
}

// Selectors represents selectors
type Selectors interface {
	List() []Selector
}

// SelectorBuilder represents a selector builder
type SelectorBuilder interface {
	Create() SelectorBuilder
	WithChain(chain chains.Chain) SelectorBuilder
	IsNot() SelectorBuilder
	Now() (Selector, error)
}

// Selector represents a selector
type Selector interface {
	Chain() chains.Chain
	IsNot() bool
}
