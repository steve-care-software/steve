package queries

import (
	"github.com/steve-care-software/steve/parsers/domain/asts"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens/elements"
)

type adapterFactory struct {
	grammarAdapter        grammars.Adapter
	astAdapter            asts.Adapter
	builder               Builder
	grammarElementBuilder elements.ElementBuilder
	chainBuilder          chains.Builder
	tokenBuilder          chains.TokenBuilder
	elementBuilder        chains.ElementBuilder
	input                 []byte
}

func createAdapterFactory(
	grammarAdapter grammars.Adapter,
	astAdapter asts.Adapter,
	builder Builder,
	grammarElementBuilder elements.ElementBuilder,
	chainBuilder chains.Builder,
	tokenBuilder chains.TokenBuilder,
	elementBuilder chains.ElementBuilder,
	input []byte,
) AdapterFactory {
	out := adapterFactory{
		grammarAdapter:        grammarAdapter,
		astAdapter:            astAdapter,
		builder:               builder,
		grammarElementBuilder: grammarElementBuilder,
		chainBuilder:          chainBuilder,
		tokenBuilder:          tokenBuilder,
		elementBuilder:        elementBuilder,
		input:                 input,
	}

	return &out
}

// Create creates a new adapter
func (app *adapterFactory) Create() (Adapter, error) {
	grammar, _, err := app.grammarAdapter.ToGrammar(app.input)
	if err != nil {
		return nil, err
	}
	return createAdapter(
		app.astAdapter,
		app.builder,
		app.grammarElementBuilder,
		app.chainBuilder,
		app.tokenBuilder,
		app.elementBuilder,
		grammar,
	), nil
}
