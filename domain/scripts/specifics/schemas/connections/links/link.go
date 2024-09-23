package links

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/schemas/points"
)

type link struct {
	hash   hash.Hash
	origin points.Point
	target points.Point
}

func createLink(
	hash hash.Hash,
	origin points.Point,
	target points.Point,
) Link {
	out := link{
		hash:   hash,
		origin: origin,
		target: target,
	}

	return &out
}

// Hash returns the hash
func (obj *link) Hash() hash.Hash {
	return obj.hash
}

// Origin returns the origin
func (obj *link) Origin() points.Point {
	return obj.origin
}

// Target returns the target
func (obj *link) Target() points.Point {
	return obj.target
}
