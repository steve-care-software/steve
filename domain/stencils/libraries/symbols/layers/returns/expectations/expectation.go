package expectations

import "github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns/kinds"

type expectation struct {
	variable string
	kind     kinds.Kind
}

func createExpectation(
	variable string,
	kind kinds.Kind,
) Expectation {
	out := expectation{
		variable: variable,
		kind:     kind,
	}

	return &out
}

// Variable returns the variable
func (obj *expectation) Variable() string {
	return obj.variable
}

// Kind returns the kind
func (obj *expectation) Kind() kinds.Kind {
	return obj.kind
}
