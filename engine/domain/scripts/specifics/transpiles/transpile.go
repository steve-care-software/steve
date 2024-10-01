package transpiles

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles/blocks"
)

type transpile struct {
	hash   hash.Hash
	blocks blocks.Blocks
	root   string
	origin string
	target string
}

func createTranspile(
	hash hash.Hash,
	blocks blocks.Blocks,
	root string,
	origin string,
	target string,
) Transpile {
	out := transpile{
		hash:   hash,
		blocks: blocks,
		root:   root,
		origin: origin,
		target: target,
	}

	return &out
}

// Hash returns the hash
func (obj *transpile) Hash() hash.Hash {
	return obj.hash
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
func (obj *transpile) Origin() string {
	return obj.origin
}

// Target returns the target
func (obj *transpile) Target() string {
	return obj.target
}
