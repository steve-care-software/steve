package walkers

import "github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"

// ElementFn represents the element func
type ElementFn func(input any) (any, error)

// MapFn takes a list of token values and returns an instance
type MapFn func(elementName string, mp map[string][]any) (any, error)

// ListFn takes a list of elements and returns an instance
type ListFn func(list []any) (any, error)

// Builder represents the walker builder
type Builder interface {
	Create() Builder
	WithFn(fn ElementFn) Builder
	WithList(list TokenList) Builder
	Now() (Walker, error)
}

// Walker represents an element walker
type Walker interface {
	Fn() ElementFn
	HasList() bool
	List() TokenList
}

// TokenListBuilder represents a token list builder
type TokenListBuilder interface {
	Create() TokenListBuilder
	WithFn(fn MapFn) TokenListBuilder
	WithList(list []SelectedTokenList) TokenListBuilder
	Now() (TokenList, error)
}

// TokenList represents a token list
type TokenList interface {
	Fn() MapFn
	List() []SelectedTokenList
}

// SelectedTokenListBuilder represents the selected token list builder
type SelectedTokenListBuilder interface {
	Create() SelectedTokenListBuilder
	WithName(name string) SelectedTokenListBuilder
	WithChain(chain chains.Chain) SelectedTokenListBuilder
	WithNode(node Node) SelectedTokenListBuilder
	Now() (SelectedTokenList, error)
}

// SelectedTokenList represents a selected token list
type SelectedTokenList interface {
	Name() string
	HasChain() bool
	Chain() chains.Chain
	HasNode() bool
	Node() Node
}

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithFn(fn ListFn) TokenBuilder
	WithNext(next Walker) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Fn() ListFn
	HasNext() bool
	Next() Walker
}

// NodeBuilder represents the node builder
type NodeBuilder interface {
	Create() NodeBuilder
	WithToken(token Token) NodeBuilder
	WithTokenList(tokenList TokenList) NodeBuilder
	WithElement(element Walker) NodeBuilder
	Now() (Node, error)
}

// Node represents a node
type Node interface {
	IsToken() bool
	Token() Token
	IsTokenList() bool
	TokenList() TokenList
	IsElement() bool
	Element() Walker
}
