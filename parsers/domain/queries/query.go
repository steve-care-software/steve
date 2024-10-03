package queries

import "github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"

type query struct {
	name    string
	version uint
	chain   chains.Chain
}

func createQuery(
	name string,
	version uint,
	chain chains.Chain,
) Query {
	out := query{
		name:    name,
		version: version,
		chain:   chain,
	}

	return &out
}

// Name returns the name
func (obj *query) Name() string {
	return obj.name
}

// Version returns the version
func (obj *query) Version() uint {
	return obj.version
}

// Chain returns the chain
func (obj *query) Chain() chains.Chain {
	return obj.chain
}
