package links

// NewLinksForTests create new links for tests
func NewLinksForTests(list []Link) Links {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkForTests creates a new link for tests
func NewLinkForTests(
	name string,
) Link {
	ins, err := NewLinkBuilder().Create().
		WithName(name).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkForTests creates a new link for tests
func NewLinkWithReverseForTests(
	name string,
	reverse string,
) Link {
	ins, err := NewLinkBuilder().Create().
		WithName(name).
		WithReverse(reverse).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
