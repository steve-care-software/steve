package symbols

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/pointers/symbols/kinds"
)

type symbol struct {
	hash hash.Hash
	name string
	kind kinds.Kind
}

func createSymbol(
	hash hash.Hash,
	name string,
	kind kinds.Kind,
) Symbol {
	out := symbol{
		hash: hash,
		name: name,
		kind: kind,
	}

	return &out
}

// Hash returns the hash
func (obj *symbol) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *symbol) Name() string {
	return obj.name
}

// Kind returns the kind
func (obj *symbol) Kind() kinds.Kind {
	return obj.kind
}
