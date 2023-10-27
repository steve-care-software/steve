package symbols

type symbol struct {
	name string
	kind uint8
}

func createSymbol(
	name string,
	kind uint8,
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
func (obj *symbol) Kind() uint8 {
	return obj.kind
}
