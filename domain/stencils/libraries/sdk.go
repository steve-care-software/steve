package libraries

import (
	"github.com/steve-care-software/steve/domain/stencils/libraries/results"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols"
)

// Library represents a library
type Library interface {
	Path() []string
	Symbols() symbols.Symbols
}

// Service represents the library service
type Service interface {
	Save(library Library) (results.Result, error)
}
