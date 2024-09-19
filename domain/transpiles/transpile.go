package transpiles

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/transpiles/blocks"
)

type transpile struct {
	blocks blocks.Blocks
	root   string
	origin hash.Hash
	target hash.Hash
}

func createTranspile(
	blocks blocks.Blocks,
	root string,
	origin hash.Hash,
	target hash.Hash,
) Transpile {
	out := transpile{
		blocks: blocks,
		root:   root,
		origin: origin,
		target: target,
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

// Origin returns the origin
func (obj *transpile) Origin() hash.Hash {
	return obj.origin
}

// Target returns the target
func (obj *transpile) Target() hash.Hash {
	return obj.target
}
