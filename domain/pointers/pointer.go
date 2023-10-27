package pointers

import "github.com/steve-care-software/steve/domain/pointers/symbols"

type pointer struct {
	path   []string
	symbol symbols.Symbol
}

func createPointer(
	path []string,
	symbol symbols.Symbol,
) Pointer {
	out := pointer{
		path:   path,
		symbol: symbol,
	}

	return &out
}

// Path returns the path
func (obj *pointer) Path() []string {
	return obj.path
}

// Symbol returns the symbol
func (obj *pointer) Symbol() symbols.Symbol {
	return obj.symbol
}
