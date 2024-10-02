package balances

import "github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the balance builder
type Builder interface {
	Create() Builder
	WithLines(lines []selectors.Selectors) Builder
	Now() (Balance, error)
}

// Balance represents balance
type Balance interface {
	Lines() []selectors.Selectors
}
