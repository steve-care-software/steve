package libraries

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols"
)

type library struct {
	hash    hash.Hash
	path    []string
	symbols symbols.Symbols
}

func createLibrary(
	hash hash.Hash,
	path []string,
	symbols symbols.Symbols,
) Library {
	out := library{
		hash:    hash,
		path:    path,
		symbols: symbols,
	}

	return &out
}

// Hash returns the hash
func (obj *library) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *library) Path() []string {
	return obj.path
}

// Symbols returns the symbols
func (obj *library) Symbols() symbols.Symbols {
	return obj.symbols
}
