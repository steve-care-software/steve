package libraries

import "github.com/steve-care-software/steve/domain/stencils/libraries/symbols"

// Library represents a library
type Library interface {
	Path() []string
	Symbols() symbols.Symbols
}
