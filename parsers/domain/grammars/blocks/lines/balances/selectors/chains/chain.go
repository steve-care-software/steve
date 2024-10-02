package chains

import "github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens/elements"

type chain struct {
	element elements.Element
	token   Token
}

func createChain(
	element elements.Element,
) Chain {
	return createChainInternally(
		element,
		nil,
	)
}

func createChainWithToken(
	element elements.Element,
	token Token,
) Chain {
	return createChainInternally(
		element,
		token,
	)
}

func createChainInternally(
	element elements.Element,
	token Token,
) Chain {
	out := chain{
		element: element,
		token:   token,
	}

	return &out
}

// Element returns the element
func (obj *chain) Element() elements.Element {
	return obj.element
}

// HasToken returns true if there is a token, false otherwise
func (obj *chain) HasToken() bool {
	return obj.token != nil
}

// Token returns the token, if any
func (obj *chain) Token() Token {
	return obj.token
}
