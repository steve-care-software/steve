package links

import "github.com/steve-care-software/steve/domain/connections/links/contexts"

type link struct {
	name     string
	isLeft   bool
	weight   float32
	contexts contexts.Contexts
}

func createLink(
	name string,
	isLeft bool,
	weight float32,
	contexts contexts.Contexts,
) Link {
	return createLinkInternally(name, isLeft, weight, contexts)
}

func createLinkInternally(
	name string,
	isLeft bool,
	weight float32,
	contexts contexts.Contexts,
) Link {
	out := link{
		name:     name,
		isLeft:   isLeft,
		weight:   weight,
		contexts: contexts,
	}

	return &out
}

// Name returns the name
func (obj *link) Name() string {
	return obj.name
}

// IsLeft returns true if left, false otherwise
func (obj *link) IsLeft() bool {
	return obj.isLeft
}

// Weight returns the weight
func (obj *link) Weight() float32 {
	return obj.weight
}

// Contexts returns the contexts, if any
func (obj *link) Contexts() contexts.Contexts {
	return obj.contexts
}
