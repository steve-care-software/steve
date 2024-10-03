package queries

import (
	"errors"

	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"
)

type builder struct {
	name    string
	version uint
	chain   chains.Chain
}

func createBuilder() Builder {
	out := builder{
		name:    "",
		version: 0,
		chain:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithVersion adds a version to the builder
func (app *builder) WithVersion(version uint) Builder {
	app.version = version
	return app
}

// WithChain adds a chain to the builder
func (app *builder) WithChain(chain chains.Chain) Builder {
	app.chain = chain
	return app
}

// Now builds a new Query instance
func (app *builder) Now() (Query, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Query instance")
	}

	if app.version <= 0 {
		return nil, errors.New("the version is mandatory in order to build a Query instance")
	}

	if app.chain == nil {
		return nil, errors.New("the chain is mandatory in order to build a Query instance")
	}

	return createQuery(
		app.name,
		app.version,
		app.chain,
	), nil
}
