package links

import "github.com/steve-care-software/steve/domain/connections/links/contexts"

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
	context contexts.Context,
	name string,
	weight float32,
) Link {
	ins, err := NewLinkBuilder().Create().
		WithName(name).
		WithContext(context).
		WithWeight(weight).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkForTests creates a new link for tests
func NewLinkWithReverseForTests(
	context contexts.Context,
	name string,
	weight float32,
	reverse string,
) Link {
	ins, err := NewLinkBuilder().Create().
		WithName(name).
		WithContext(context).
		WithWeight(weight).
		WithReverse(reverse).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
