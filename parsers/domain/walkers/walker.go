package walkers

type walker struct {
	fn   ElementFn
	list TokenList
}

func createWalker(
	fn ElementFn,
) Walker {
	return createWalkerInternally(fn, nil)
}

func createWalkerWithList(
	fn ElementFn,
	list TokenList,
) Walker {
	return createWalkerInternally(fn, list)
}

func createWalkerInternally(
	fn ElementFn,
	list TokenList,
) Walker {
	out := walker{
		fn:   fn,
		list: list,
	}

	return &out
}

// Fn returns the ElementFn
func (obj *walker) Fn() ElementFn {
	return obj.fn
}

// Fn returns the ElementFn
func (obj *walker) HasList() bool {
	return obj.list != nil
}

// List returns the list
func (obj *walker) List() TokenList {
	return obj.list
}
