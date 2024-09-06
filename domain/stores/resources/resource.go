package resources

import "github.com/steve-care-software/steve/domain/stores/resources/pointers"

type resource struct {
	identifier string
	pointers   pointers.Pointers
}

func createResource(
	identifier string,
	pointers pointers.Pointers,
) Resource {
	out := resource{
		identifier: identifier,
		pointers:   pointers,
	}

	return &out
}

// Identifier returns the identifier
func (obj *resource) Identifier() string {
	return obj.identifier
}

// Pointers returns the pointers
func (obj *resource) Pointers() pointers.Pointers {
	return obj.pointers
}
