package walkers

type node struct {
	token     Token
	tokenList TokenList
	element   Walker
}

func createNodeWithToken(
	token Token,
) Node {
	return createNodeInternally(token, nil, nil)
}

func createNodeWithTokenList(
	tokenList TokenList,
) Node {
	return createNodeInternally(nil, tokenList, nil)
}

func createNodeWithElement(
	element Walker,
) Node {
	return createNodeInternally(nil, nil, element)
}

func createNodeInternally(
	token Token,
	tokenList TokenList,
	element Walker,
) Node {
	out := node{
		token:     token,
		tokenList: tokenList,
		element:   element,
	}

	return &out
}

// IsToken returns true if there is a token, false otherwise
func (obj *node) IsToken() bool {
	return obj.token != nil
}

// Token returns the token, if any
func (obj *node) Token() Token {
	return obj.token
}

// IsTokenList returns true if there is a token list, false otherwise
func (obj *node) IsTokenList() bool {
	return obj.tokenList != nil
}

// TokenList returns the token list, if any
func (obj *node) TokenList() TokenList {
	return obj.tokenList
}

// IsElement returns true if there is an element, false otherwise
func (obj *node) IsElement() bool {
	return obj.element != nil
}

// Element returns the element, if any
func (obj *node) Element() Walker {
	return obj.element
}
