package lines

import (
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewLineBuilder creates a line builder
func NewLineBuilder() LineBuilder {
	return createLineBuilder()
}

// Builder represents a line builder
type Builder interface {
	Create() Builder
	WithList(list []Line) Builder
	Now() (Lines, error)
}

// Lines represents lines
type Lines interface {
	List() []Line
}

// LineBuilder represents a line builder
type LineBuilder interface {
	Create() LineBuilder
	WithTokens(tokens tokens.Tokens) LineBuilder
	WithBalance(balance balances.Balance) LineBuilder
	Now() (Line, error)
}

// Line represents a variable
type Line interface {
	Tokens() tokens.Tokens
	HasBalance() bool
	Balance() balances.Balance
}
