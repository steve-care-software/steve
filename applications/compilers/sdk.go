package compilers

import "github.com/steve-care-software/steve/domain/stencils/libraries/symbols"

// Application represents a compiler application
type Application interface {
	Execute(input []byte) (symbols.Symbols, error)
}
