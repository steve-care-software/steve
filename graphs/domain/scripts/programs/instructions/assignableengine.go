package instructions

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/grammars"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries"
	selectors_chain "github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"
)

type assignableEngine struct {
	selector selectors_chain.Chain
	grammar  grammars.Grammar
	query    queries.Query
}

func createAssignableEngineWithSelector(selector selectors_chain.Chain) AssignableEngine {
	return createAssignableEngineInternally(selector, nil, nil)
}

// createAssignableEngineWithGrammar creates an assignable engine with a grammar
func createAssignableEngineWithGrammar(grammar grammars.Grammar) AssignableEngine {
	return createAssignableEngineInternally(nil, grammar, nil)
}

// createAssignableEngineWithQuery creates an assignable engine with a query
func createAssignableEngineWithQuery(query queries.Query) AssignableEngine {
	return createAssignableEngineInternally(nil, nil, query)
}

func createAssignableEngineInternally(
	selector selectors_chain.Chain,
	grammar grammars.Grammar,
	query queries.Query,
) AssignableEngine {
	out := assignableEngine{
		selector: selector,
		grammar:  grammar,
		query:    query,
	}

	return &out
}

// IsSelector returns true if the engine is a selector
func (obj *assignableEngine) IsSelector() bool {
	return obj.selector != nil
}

// Selector returns the selector if present
func (obj *assignableEngine) Selector() selectors_chain.Chain {
	return obj.selector
}

// IsGrammar returns true if the engine is a grammar
func (obj *assignableEngine) IsGrammar() bool {
	return obj.grammar != nil
}

// Grammar returns the grammar if present
func (obj *assignableEngine) Grammar() grammars.Grammar {
	return obj.grammar
}

// IsQuery returns true if the engine is a query
func (obj *assignableEngine) IsQuery() bool {
	return obj.query != nil
}

// Query returns the query if present
func (obj *assignableEngine) Query() queries.Query {
	return obj.query
}
