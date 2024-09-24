package blocks

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/programs/grammars/blocks/suites"
	"github.com/steve-care-software/steve/domain/scripts/specifics/grammars/blocks/lines"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewBlockBuilder creates a new block builder
func NewBlockBuilder() BlockBuilder {
	hashAdapter := hash.NewAdapter()
	return createBlockBuilder(
		hashAdapter,
	)
}

// Builder represents a blocks builder
type Builder interface {
	Create() Builder
	WithList(list []Block) Builder
	Now() (Blocks, error)
}

// Blocks represents blocks
type Blocks interface {
	Hash() hash.Hash
	List() []Block
}

// BlockBuilder represents the block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithName(name string) BlockBuilder
	WithLines(lines lines.Lines) BlockBuilder
	WithSuites(suites suites.Suites) BlockBuilder
	Now() (Block, error)
}

// Block represents a block
type Block interface {
	Hash() hash.Hash
	Name() string
	Lines() lines.Lines
	HasSuites() bool
	Suites() suites.Suites
}
