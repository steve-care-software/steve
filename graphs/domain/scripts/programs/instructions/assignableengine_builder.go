package instructions

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/grammars"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries"
	selectors_chain "github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"
)

type assignableEngineBuilder struct {
	selector selectors_chain.Chain
	grammar  grammars.Grammar
	query    queries.Query
}

func createAssignableEngineBuilder() AssignableEngineBuilder {
	return &assignableEngineBuilder{
		selector: nil,
		grammar:  nil,
		query:    nil,
	}
}

// Create initializes the assignable engine builder
func (obj *assignableEngineBuilder) Create() AssignableEngineBuilder {
	return createAssignableEngineBuilder()
}

// WithSelector adds a selector to the builder
func (obj *assignableEngineBuilder) WithSelector(selector selectors_chain.Chain) AssignableEngineBuilder {
	obj.selector = selector
	return obj
}

// WithGrammar adds a grammar to the builder
func (obj *assignableEngineBuilder) WithGrammar(grammar grammars.Grammar) AssignableEngineBuilder {
	obj.grammar = grammar
	return obj
}

// WithQuery adds a query to the builder
func (obj *assignableEngineBuilder) WithQuery(query queries.Query) AssignableEngineBuilder {
	obj.query = query
	return obj
}

// Now builds a new AssignableEngine instance
func (obj *assignableEngineBuilder) Now() (AssignableEngine, error) {
	if obj.selector != nil {
		return createAssignableEngineWithSelector(obj.selector), nil
	}

	if obj.grammar != nil {
		return createAssignableEngineWithGrammar(obj.grammar), nil
	}

	if obj.query != nil {
		return createAssignableEngineWithQuery(obj.query), nil
	}

	return nil, errors.New("the AssignableEngine is invalid: a selector, grammar, or query must be provided")
}
