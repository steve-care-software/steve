package blocks

import (
	"github.com/steve-care-software/steve/engine/domain/blockchains/blocks/contents"
)

// NewBlocksForTests creates a new blocks for tests
func NewBlocksForTests(list []Block) Blocks {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewBlockForTests creates a new block for tests
func NewBlockForTests(content contents.Content, result []byte) Block {
	ins, err := NewBlockBuilder().Create().WithContent(content).WithResult(result).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
