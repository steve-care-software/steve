package balances

import "github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors"

type balance struct {
	lines []selectors.Selectors
}

func createBalance(
	lines []selectors.Selectors,
) Balance {
	out := balance{
		lines: lines,
	}

	return &out
}

// Lines returns the lines
func (obj *balance) Lines() []selectors.Selectors {
	return obj.lines
}
