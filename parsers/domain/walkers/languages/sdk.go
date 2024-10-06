package languages

import (
	"github.com/steve-care-software/steve/parsers/domain/walkers"
)

// Adapter represents an element adapter
type Adapter interface {
	ToWalker(ins Element) (walkers.Walker, error)
}

// Element represents an element
type Element struct {
	ElementFn walkers.ElementFn
	TokenList *TokenList
}

// TokenList represents the token list
type TokenList struct {
	List  map[string]SelectedTokenList
	MapFn walkers.MapFn
}

// SelectedTokenList represents the selected token list
type SelectedTokenList struct {
	SelectorScript []byte
	Node           *Node
}

// Token represents a token
type Token struct {
	ListFn walkers.ListFn
	Next   *Element
}

// Node represents a node
type Node struct {
	Token     *Token
	TokenList *TokenList
	Element   *Element
}
