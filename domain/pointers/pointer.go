package pointers

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/pointers/symbols"
)

type pointer struct {
	hash   hash.Hash
	path   []string
	symbol symbols.Symbol
}

func createPointer(
	hash hash.Hash,
	path []string,
	symbol symbols.Symbol,
) Pointer {
	out := pointer{
		hash:   hash,
		path:   path,
		symbol: symbol,
	}

	return &out
}

// Hash returns the hash
func (obj *pointer) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *pointer) Path() []string {
	return obj.path
}

// Symbol returns the symbol
func (obj *pointer) Symbol() symbols.Symbol {
	return obj.symbol
}
