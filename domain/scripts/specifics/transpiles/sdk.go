package transpiles

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/transpiles/blocks"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// FetchGrammarInput returns the grammar input
func FetchGrammarInput() []byte {
	return grammarInput()
}

// ToTranspile converts an input to a transpile instance
type ParserAdapter interface {
	// ToTranspile takes the input and create a transpile instance
	ToTranspile(input []byte) (Transpile, []byte, error)
}

// Builder represents the transpile builder
type Builder interface {
	Create() Builder
	WithBlocks(blocks blocks.Blocks) Builder
	WithRoot(root string) Builder
	WithOrigin(origin string) Builder
	WithTarget(target string) Builder
	Now() (Transpile, error)
}

// Transpile represents a transpile
type Transpile interface {
	Hash() hash.Hash
	Blocks() blocks.Blocks
	Root() string
	Origin() string
	Target() string
}
