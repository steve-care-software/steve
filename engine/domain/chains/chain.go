package chains

import (
	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/suites"
)

type chain struct {
	grammar hash.Hash
	action  Action
	suites  suites.Suites
}

func createChain(
	grammar hash.Hash,
	action Action,
) Chain {
	return createChainInternally(grammar, action, nil)
}

func createChainWithSuites(
	grammar hash.Hash,
	action Action,
	suites suites.Suites,
) Chain {
	return createChainInternally(grammar, action, suites)
}

func createChainInternally(
	grammar hash.Hash,
	action Action,
	suites suites.Suites,
) Chain {
	out := chain{
		grammar: grammar,
		action:  action,
		suites:  suites,
	}

	return &out
}

// Grammar returns the grammar
func (obj *chain) Grammar() hash.Hash {
	return obj.grammar
}

// Action returns the action
func (obj *chain) Action() Action {
	return obj.action
}

// HasSuites returns true if there is suites, false otherwise
func (obj *chain) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *chain) Suites() suites.Suites {
	return obj.suites
}
