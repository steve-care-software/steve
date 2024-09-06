package resources

import "github.com/steve-care-software/steve/domain/stores/resources/pointers"

type resource struct {
	identifier string
	pointers   pointers.Pointer
}

func createResource(
	identifier string,
	pointers pointers.Pointer,
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

// Pointer returns the pointers
func (obj *resource) Pointer() pointers.Pointer {
	return obj.pointers
}
