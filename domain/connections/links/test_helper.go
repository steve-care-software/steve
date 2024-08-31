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
	isLeft bool,
	weight float32,
) Link {
	builder := NewLinkBuilder().Create().
		WithName(name).
		WithContext(context).
		WithWeight(weight)

	if isLeft {
		builder.IsLeft()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
