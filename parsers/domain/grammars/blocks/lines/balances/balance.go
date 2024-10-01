package balances

import "github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/operations"

type balance struct {
	lines []operations.Operations
}

func createBalance(
	lines []operations.Operations,
) Balance {
	out := balance{
		lines: lines,
	}

	return &out
}

// Lines returns the lines
func (obj *balance) Lines() []operations.Operations {
	return obj.lines
}
