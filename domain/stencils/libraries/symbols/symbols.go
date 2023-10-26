package symbols

type symbols struct {
	list []Symbol
}

func createSymbols(
	list []Symbol,
) Symbols {
	out := symbols{
		list: list,
	}

	return &out
}

// List returns the symbols list
func (obj *symbols) List() []Symbol {
	return obj.list
}

// Fetch fetches a symbol by name
func (obj *symbols) Fetch(name string) (Symbol, error) {
	return nil, nil
}
