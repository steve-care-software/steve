package walkers

type tokenlist struct {
	fn   MapFn
	list []SelectedTokenList
}

func createTokenList(
	fn MapFn,
	list []SelectedTokenList,
) TokenList {
	out := tokenlist{
		fn:   fn,
		list: list,
	}

	return &out
}

// Fn returns the MapFn
func (obj *tokenlist) Fn() MapFn {
	return obj.fn
}

// List returns the list
func (obj *tokenlist) List() []SelectedTokenList {
	return obj.list
}
