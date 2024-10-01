package blocks

import (
	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles/blocks/lines"
)

type block struct {
	hash  hash.Hash
	name  string
	lines lines.Lines
}

func createBlock(
	hash hash.Hash,
	name string,
	lines lines.Lines,
) Block {
	out := block{
		hash:  hash,
		name:  name,
		lines: lines,
	}

	return &out
}

// Hash returns the hash
func (obj *block) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *block) Name() string {
	return obj.name
}

// Lines returns the lines
func (obj *block) Lines() lines.Lines {
	return obj.lines
}
