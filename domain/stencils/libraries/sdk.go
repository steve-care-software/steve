package libraries

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/results"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a library builder
type Builder interface {
	Create() Builder
	WithPath(path []string) Builder
	WithSybols(symbols symbols.Symbols) Builder
	Now() (Library, error)
}

// Library represents a library
type Library interface {
	Hash() hash.Hash
	Path() []string
	Symbols() symbols.Symbols
}

// Service represents the library service
type Service interface {
	Save(library Library) (results.Result, error)
}
