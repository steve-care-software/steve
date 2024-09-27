package lines

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/programs/grammars/blocks/suites"
	"github.com/steve-care-software/steve/domain/scripts/specifics/transpiles/blocks/lines/tokens"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewLineBuilder creates a new line builder
func NewLineBuilder() LineBuilder {
	hashAdapter := hash.NewAdapter()
	return createLineBuilder(
		hashAdapter,
	)
}

// Builder represents lines builder
type Builder interface {
	Create() Builder
	WithList(list []Line) Builder
	Now() (Lines, error)
}

// Lines represents lines
type Lines interface {
	Hash() hash.Hash
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
	Hash() hash.Hash
	Tokens() tokens.Tokens
	HasSuites() bool
	Suites() suites.Suites
}
