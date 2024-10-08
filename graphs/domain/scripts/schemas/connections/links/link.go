package links

import "github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references"

type link struct {
	origin references.Reference
	target references.Reference
}

func createLink(
	origin references.Reference,
	target references.Reference,
) Link {
	out := link{
		origin: origin,
		target: target,
	}

	return &out
}

// Origin returns the origin
func (obj *link) Origin() references.Reference {
	return obj.origin
}

// Target returns the target
func (obj *link) Target() references.Reference {
	return obj.target
}
