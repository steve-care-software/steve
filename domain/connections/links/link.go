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
	isLeft  bool
	weight  float32
}

func createLink(
	hash hash.Hash,
	context contexts.Context,
	name string,
	isLeft bool,
	weight float32,
) Link {
	out := link{
		hash:    hash,
		context: context,
		name:    name,
		isLeft:  isLeft,
		weight:  weight,
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

// IsLeft returns true if left, false otherwise
func (obj *link) IsLeft() bool {
	return obj.isLeft
}

// Weight returns the weight
func (obj *link) Weight() float32 {
	return obj.weight
}

// Debug returns the string debug representation of the link
func (obj *link) Debug() string {
	return fmt.Sprintf("name: %s, isLeft: %t, weight: %f", obj.name, obj.isLeft, obj.weight)
}
