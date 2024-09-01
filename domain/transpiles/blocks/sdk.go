package blocks

import "github.com/steve-care-software/steve/domain/transpiles/blocks/lines"

// Blocks represents blocks
type Blocks interface {
	List() []Block
}

// Block represents a block
type Block interface {
	Name() string
	Lines() lines.Lines
}
