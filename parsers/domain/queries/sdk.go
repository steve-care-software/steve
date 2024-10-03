package queries

import (
	"github.com/steve-care-software/steve/parsers/domain/asts"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens/elements"
)

// NewAdapterFactory creates a new adapter factory
func NewAdapterFactory() AdapterFactory {
	grammarAdapter := grammars.NewAdapter()
	astAdapter := asts.NewAdapter()
	builder := NewBuilder()
	grammarElementBuilder := elements.NewElementBuilder()
	chainBuilder := chains.NewBuilder()
	tokenBuilder := chains.NewTokenBuilder()
	elementBuilder := chains.NewElementBuilder()
	input := fetchGrammarInput()
	return createAdapterFactory(
		grammarAdapter,
		astAdapter,
		builder,
		grammarElementBuilder,
		chainBuilder,
		tokenBuilder,
		elementBuilder,
		input,
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// AdapterFactory represents an adapter factory
type AdapterFactory interface {
	Create() (Adapter, error)
}

// Adapter represents an adapter
type Adapter interface {
	ToQuery(input []byte) (Query, []byte, error)
}

// Builder represents the query builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithVersion(version uint) Builder
	WithChain(chain chains.Chain) Builder
	Now() (Query, error)
}

// Query represents a query
type Query interface {
	Name() string
	Version() uint
	Chain() chains.Chain
}
