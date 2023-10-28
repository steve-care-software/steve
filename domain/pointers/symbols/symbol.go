package symbols

import "github.com/steve-care-software/steve/domain/pointers/symbols/kinds"

type symbol struct {
	name string
	kind kinds.Kind
}

func createSymbol(
	name string,
	kind kinds.Kind,
) Symbol {
	out := symbol{
		name: name,
		kind: kind,
	}

	return &out
}

// Name returns the name
func (obj *symbol) Name() string {
	return obj.name
}

// Kind returns the kind
func (obj *symbol) Kind() kinds.Kind {
	return obj.kind
}
