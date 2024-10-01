package chains

import "github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens/elements"

type chain struct {
	element      elements.Element
	tokenIndex   uint
	elementIndex uint
	next         Chain
}

func createChain(
	element elements.Element,
	tokenIndex uint,
	elementIndex uint,
) Chain {
	return createChainInternally(
		element,
		tokenIndex,
		elementIndex,
		nil,
	)
}

func createChainWithNext(
	element elements.Element,
	tokenIndex uint,
	elementIndex uint,
	next Chain,
) Chain {
	return createChainInternally(
		element,
		tokenIndex,
		elementIndex,
		next,
	)
}

func createChainInternally(
	element elements.Element,
	tokenIndex uint,
	elementIndex uint,
	next Chain,
) Chain {
	out := chain{
		element:      element,
		tokenIndex:   tokenIndex,
		elementIndex: elementIndex,
		next:         next,
	}

	return &out
}

// Element returns the element
func (obj *chain) Element() elements.Element {
	return obj.element
}

// TokenIndex returns the tokenIndex
func (obj *chain) TokenIndex() uint {
	return obj.tokenIndex
}

// ElementIndex returns the elementIndex
func (obj *chain) ElementIndex() uint {
	return obj.elementIndex
}

// HasNext returns true if there is a next, false otherwise
func (obj *chain) HasNext() bool {
	return obj.next != nil
}

// Next returns the next, if any
func (obj *chain) Next() Chain {
	return obj.next
}
