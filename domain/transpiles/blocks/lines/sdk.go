package lines

import (
	"github.com/steve-care-software/steve/domain/programs/grammars/blocks/suites"
	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens"
)

// Builder represents lines builder
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
	WithSuites(suites suites.Suites) LineBuilder
	Now() (Line, error)
}

// Line represents a line
type Line interface {
	Tokens() tokens.Tokens
	HasSuites() bool
	Suites() suites.Suites
}
