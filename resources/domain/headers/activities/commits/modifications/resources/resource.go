package resources

import (
	"github.com/steve-care-software/steve/hash"
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits/modifications/resources/pointers"
)

type resource struct {
	hash       hash.Hash
	identifier string
	pointer    pointers.Pointer
}

func createResource(
	hash hash.Hash,
	identifier string,
	pointer pointers.Pointer,
) Resource {
	out := resource{
		hash:       hash,
		identifier: identifier,
		pointer:    pointer,
	}

	return &out
}

// Hash returns the hash
func (obj *resource) Hash() hash.Hash {
	return obj.hash
}

// Identifier returns the identifier
func (obj *resource) Identifier() string {
	return obj.identifier
}

// Pointer returns the pointer
func (obj *resource) Pointer() pointers.Pointer {
	return obj.pointer
}
