package balances

import "github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/operations"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the balance builder
type Builder interface {
	Create() Builder
	WithLines(lines []operations.Operations) Builder
	Now() (Balance, error)
}

// Balance represents balance
type Balance interface {
	Lines() []operations.Operations
}
