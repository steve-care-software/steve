package applications

import (
	"github.com/steve-care-software/steve/parsers/domain/asts"
	"github.com/steve-care-software/steve/parsers/domain/asts/instructions"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
	"github.com/steve-care-software/steve/parsers/domain/queries"
)

// ElementFn represents the element func
type ElementFn func(input any) (any, error)

// MapFn takes a list of token values and returns an instance
type MapFn func(elementName string, mp map[string][]any) (any, error)

// ListFn takes a list of elements and returns an instance
type ListFn func(list []any) (any, error)

// NewApplication creates a new application
func NewApplication() Application {
	elementsAdapter := instructions.NewElementsAdapter()
	astAdapter := asts.NewAdapter()
	queryAdapter, _ := queries.NewAdapterFactory().Create()
	tokensBuilder := instructions.NewTokensBuilder()
	return createApplication(
		elementsAdapter,
		astAdapter,
		queryAdapter,
		tokensBuilder,
	)
}

// Element represents an element
type Element struct {
	ElementFn ElementFn
	TokenList *TokenList
}

// TokenList represents the token list
type TokenList struct {
	List  map[string]ChosenTokenList
	MapFn MapFn
}

// ChosenTokenList represents the chosen token list
type ChosenTokenList struct {
	SelectorScript []byte
	Node           *Node
}

// Token represents a token
type Token struct {
	ListFn ListFn
	Next   *Element
}

// Node represents a node
type Node struct {
	Token     *Token
	TokenList *TokenList
	Element   *Element
}

// Application represents the interpreter application
type Application interface {
	// Execute executes the parser
	Execute(input []byte, grammar grammars.Grammar, ins Element) (any, []byte, error)

	// Suites executes all the test suites of the grammar
	Suites(grammar grammars.Grammar) error
}
