package selectors

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/operations/selectors/chains"
)

// Builder represents a selector builder
type Builder interface {
	Create() Builder
	WithChain(chain chains.Chain) Builder
	IsNot() Builder
	Now() (Selector, error)
}

// Selector represents a selector
type Selector interface {
	Hash() hash.Hash
	Chain() chains.Chain
	IsNot() bool
}
