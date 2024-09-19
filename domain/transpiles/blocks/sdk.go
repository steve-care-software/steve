package blocks

import (
	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewBlockBuilder creates a new block builder
func NewBlockBuilder() BlockBuilder {
	return createBlockBuilder()
}

// Builder represents a blocks builder
type Builder interface {
	Create() Builder
	WithList(list []Block) Builder
	Now() (Blocks, error)
}

// Blocks represents blocks
type Blocks interface {
	List() []Block
}

// BlockBuilder represents the block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithName(name string) BlockBuilder
	WithLines(lines lines.Lines) BlockBuilder
	Now() (Block, error)
}

// Block represents a block
type Block interface {
	Name() string
	Lines() lines.Lines
}
