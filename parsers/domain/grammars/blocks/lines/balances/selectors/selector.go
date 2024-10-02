package selectors

import "github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"

type selector struct {
	chain chains.Chain
	isNot bool
}

func createSelector(
	chain chains.Chain,
	isNot bool,
) Selector {
	out := selector{
		chain: chain,
		isNot: isNot,
	}

	return &out
}

// Chain returns the chain
func (obj *selector) Chain() chains.Chain {
	return obj.chain
}

// IsNot returns true if not, false otherwise
func (obj *selector) IsNot() bool {
	return obj.isNot
}
