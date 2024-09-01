package transpiles

import (
	"github.com/steve-care-software/steve/domain/programs/grammars"
	"github.com/steve-care-software/steve/domain/transpiles/blocks"
)

// Transpile represents a transpile
type Transpile interface {
	Origin() grammars.Grammar
	Target() grammars.Grammar
	Blocks() blocks.Blocks
	Root() string
}
