package links

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/schemas/points"
)

// Links represents links
type Links interface {
	Hash() hash.Hash
	List() []Link
}

// Link represents a point link
type Link interface {
	Hash() hash.Hash
	Origin() points.Point
	Target() points.Point
}
