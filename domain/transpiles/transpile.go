package transpiles

import "github.com/steve-care-software/steve/domain/transpiles/blocks"

type transpile struct {
	blocks blocks.Blocks
	root   string
}

func createTranspile(
	blocks blocks.Blocks,
	root string,
) Transpile {
	out := transpile{
		blocks: blocks,
		root:   root,
	}

	return &out
}

// Blocks returns the blocks
func (obj *transpile) Blocks() blocks.Blocks {
	return obj.blocks
}

// Root returns the root
func (obj *transpile) Root() string {
	return obj.root
}
