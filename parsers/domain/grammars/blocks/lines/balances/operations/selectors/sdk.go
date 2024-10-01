package selectors

import (
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/operations/selectors/chains"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a selector builder
type Builder interface {
	Create() Builder
	WithChain(chain chains.Chain) Builder
	IsNot() Builder
	Now() (Selector, error)
}

// Selector represents a selector
type Selector interface {
	Chain() chains.Chain
	IsNot() bool
}
