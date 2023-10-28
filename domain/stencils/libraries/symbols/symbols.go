package symbols

import "github.com/steve-care-software/steve/domain/hash"

type symbols struct {
	hash hash.Hash
	list []Symbol
}

func createSymbols(
	hash hash.Hash,
	list []Symbol,
) Symbols {
	out := symbols{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *symbols) Hash() hash.Hash {
	return obj.hash
}

// List returns the symbols list
func (obj *symbols) List() []Symbol {
	return obj.list
}

// Fetch fetches a symbol by name
func (obj *symbols) Fetch(name string) (Symbol, error) {
	return nil, nil
}
