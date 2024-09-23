package blocks

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/programs/grammars/blocks/suites"
	"github.com/steve-care-software/steve/domain/scripts/specifics/grammars/blocks/lines"
)

// Blocks represents blocks
type Blocks interface {
	Hash() hash.Hash
	List() []Block
}

// Block represents a block
type Block interface {
	Hash() hash.Hash
	Name() string
	Lines() lines.Lines
	HasSuites() bool
	Suites() suites.Suites
}
