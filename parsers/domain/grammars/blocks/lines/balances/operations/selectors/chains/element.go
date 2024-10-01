package chains

type element struct {
	index uint
	chain Chain
}

func createElement(
	index uint,
) Element {
	return createElementInternally(index, nil)
}

func createElementWithChain(
	index uint,
	chain Chain,
) Element {
	return createElementInternally(index, chain)
}

func createElementInternally(
	index uint,
	chain Chain,
) Element {
	out := element{
		index: index,
		chain: chain,
	}

	return &out
}

// Index returns the index
func (obj *element) Index() uint {
	return obj.index
}

// HasChain returns true if there is a chain, false otherwise
func (obj *element) HasChain() bool {
	return obj.chain != nil
}

// Chain returns the chain, if any
func (obj *element) Chain() Chain {
	return obj.chain
}
