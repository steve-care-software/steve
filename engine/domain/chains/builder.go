package chains

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/suites"
)

type builder struct {
	grammar hash.Hash
	action  Action
	suites  suites.Suites
}

func createBuilder() Builder {
	out := builder{
		grammar: nil,
		action:  nil,
		suites:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithGrammar adds a grammar to the builder
func (app *builder) WithGrammar(grammar hash.Hash) Builder {
	app.grammar = grammar
	return app
}

// WithAction adds an action to the builder
func (app *builder) WithAction(action Action) Builder {
	app.action = action
	return app
}

// WithSuites add suites to the builder
func (app *builder) WithSuites(suites suites.Suites) Builder {
	app.suites = suites
	return app
}

// Now builds a new Chain instance
func (app *builder) Now() (Chain, error) {
	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build a Chain instance")
	}

	if app.action == nil {
		return nil, errors.New("the action is mandatory in order to build a Chain instance")
	}

	if app.suites != nil {
		return createChainWithSuites(app.grammar, app.action, app.suites), nil
	}

	return createChain(app.grammar, app.action), nil
}
