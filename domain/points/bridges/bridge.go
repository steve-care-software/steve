package bridges

import (
	"github.com/steve-care-software/steve/domain/connections/links"
	"github.com/steve-care-software/steve/domain/hash"
)

type bridge struct {
	hash   hash.Hash
	link   links.Link
	weight float32
}

func createBridge(
	hash hash.Hash,
	link links.Link,
	weight float32,
) Bridge {
	out := bridge{
		hash:   hash,
		link:   link,
		weight: weight,
	}

	return &out
}

// Hash returns the hash
func (obj *bridge) Hash() hash.Hash {
	return obj.hash
}

// Link returns the link
func (obj *bridge) Link() links.Link {
	return obj.link
}

// Weight returns the weight
func (obj *bridge) Weight() float32 {
	return obj.weight
}
