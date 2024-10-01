package blocks

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/grammars/blocks/lines"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/suites"
)

type block struct {
	hash   hash.Hash
	name   string
	lines  lines.Lines
	suites suites.Suites
}

func createBlock(
	hash hash.Hash,
	name string,
	lines lines.Lines,
) Block {
	return createBlockInternally(
		hash,
		name,
		lines,
		nil,
	)
}

func createBlockWithSuites(
	hash hash.Hash,
	name string,
	lines lines.Lines,
	suites suites.Suites,
) Block {
	return createBlockInternally(
		hash,
		name,
		lines,
		suites,
	)
}

func createBlockInternally(
	hash hash.Hash,
	name string,
	lines lines.Lines,
	suites suites.Suites,
) Block {
	out := block{
		hash:   hash,
		name:   name,
		lines:  lines,
		suites: suites,
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

// HasSuites returns true if there is suites, false otherwise
func (obj *block) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *block) Suites() suites.Suites {
	return obj.suites
}
