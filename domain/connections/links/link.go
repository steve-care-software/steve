package links

import (
	"fmt"

	"github.com/steve-care-software/steve/domain/connections/links/contexts"
	"github.com/steve-care-software/steve/domain/hash"
)

type link struct {
	hash    hash.Hash
	context contexts.Context
	name    string
	weight  float32
	reverse string
}

func createLink(
	hash hash.Hash,
	context contexts.Context,
	name string,
	weight float32,
) Link {
	return createLinkInternally(hash, context, name, weight, "")
}

func createLinkWithReverse(
	hash hash.Hash,
	context contexts.Context,
	name string,
	weight float32,
	reverse string,
) Link {
	return createLinkInternally(hash, context, name, weight, reverse)
}

func createLinkInternally(
	hash hash.Hash,
	context contexts.Context,
	name string,
	weight float32,
	reverse string,
) Link {
	out := link{
		hash:    hash,
		context: context,
		name:    name,
		weight:  weight,
		reverse: reverse,
	}

	return &out
}

// Hash returns the hash, if any
func (obj *link) Hash() hash.Hash {
	return obj.hash
}

// Contexts returns the context, if any
func (obj *link) Context() contexts.Context {
	return obj.context
}

// Name returns the name
func (obj *link) Name() string {
	return obj.name
}

// Weight returns the weight
func (obj *link) Weight() float32 {
	return obj.weight
}

// HasReverse returns true if there is a reverse, false otherwise
func (obj *link) HasReverse() bool {
	return obj.reverse != ""
}

// Reverse returns the reverse, if any
func (obj *link) Reverse() string {
	return obj.reverse
}

// Debug returns the string debug representation of the link
func (obj *link) Debug() string {
	reverseStr := ""
	if obj.HasReverse() {
		reverseStr = fmt.Sprintf("reverse: %s, ", obj.Reverse())
	}
	return fmt.Sprintf("name: %s, %sweight: %f", obj.name, reverseStr, obj.weight)
}
