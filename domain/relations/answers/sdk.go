package answers

import (
	"github.com/steve-care-software/steve/domain/relations/data/connections/links"
	"github.com/steve-care-software/steve/domain/relations/data/points"
)

// Answer represents an answer
type Answer interface {
	IsPoint() bool
	Point() points.Point
	IsPoints() bool
	Points() points.Points
	IsLink() bool
	Link() links.Links
	IsLinks() bool
	Links() links.Links
}
