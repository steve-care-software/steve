package links

import (
	"github.com/steve-care-software/steve/domain/connections/links/contexts"
	"github.com/steve-care-software/steve/domain/hash"
)

type link struct {
	hash     hash.Hash
	contexts contexts.Contexts
	name     string
	isLeft   bool
	weight   float32
}

func createLink(
	hash hash.Hash,
	contexts contexts.Contexts,
	name string,
	isLeft bool,
	weight float32,
) Link {
	out := link{
		hash:     hash,
		contexts: contexts,
		name:     name,
		isLeft:   isLeft,
		weight:   weight,
	}

	return &out
}

// Hash returns the hash, if any
func (obj *link) Hash() hash.Hash {
	return obj.hash
}

// Contexts returns the contexts, if any
func (obj *link) Contexts() contexts.Contexts {
	return obj.contexts
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
