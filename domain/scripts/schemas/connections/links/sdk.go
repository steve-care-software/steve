package links

import "github.com/steve-care-software/steve/domain/scripts/schemas/points"

// Links represents links
type Links interface {
	List() []Link
}

// Link represents a point link
type Link interface {
	Origin() points.Point
	Target() points.Point
}
