package walkers

import "github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"

type selectedTokenList struct {
	name  string
	chain chains.Chain
	node  Node
}

func createSelectedTokenList(
	name string,
) SelectedTokenList {
	return createSelectedTokenListInternally(name, nil, nil)
}

func createSelectedTokenListWithChain(
	name string,
	chain chains.Chain,
) SelectedTokenList {
	return createSelectedTokenListInternally(name, chain, nil)
}

func createSelectedTokenListWithNode(
	name string,
	node Node,
) SelectedTokenList {
	return createSelectedTokenListInternally(name, nil, node)
}

func createSelectedTokenListWithChainAndNode(
	name string,
	chain chains.Chain,
	node Node,
) SelectedTokenList {
	return createSelectedTokenListInternally(name, chain, node)
}

func createSelectedTokenListInternally(
	name string,
	chain chains.Chain,
	node Node,
) SelectedTokenList {
	out := selectedTokenList{
		name:  name,
		chain: chain,
		node:  node,
	}

	return &out
}

// Name returns the name
func (obj *selectedTokenList) Name() string {
	return obj.name
}

// HasChain returns true if there is a chain, false otherwise
func (obj *selectedTokenList) HasChain() bool {
	return obj.chain != nil
}

// Chain returns the chain, if any
func (obj *selectedTokenList) Chain() chains.Chain {
	return obj.chain
}

// HasNode returns true if there is a node, false otherwise
func (obj *selectedTokenList) HasNode() bool {
	return obj.node != nil
}

// Node returns the node, if any
func (obj *selectedTokenList) Node() Node {
	return obj.node
}
