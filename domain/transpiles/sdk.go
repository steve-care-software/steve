package transpiles

import (
	"github.com/steve-care-software/steve/domain/transpiles/blocks"
)

// Transpile represents a transpile
type Transpile interface {
	Blocks() blocks.Blocks
	Root() string
}
