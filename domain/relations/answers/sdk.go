package answers

import (
	"github.com/steve-care-software/steve/domain/relations/data/connections/links"
	"github.com/steve-care-software/steve/domain/relations/data/points"
)

// Builder represents the answer builder
type Builder interface {
	Create() Builder
	WithPoint(point points.Point) Builder
	WithPoints(points points.Points) Builder
	WithLink(link links.Link) Builder
	WithLinks(links links.Links) Builder
	Now() (Answer, error)
}

// Answer represents an answer
type Answer interface {
	IsPoint() bool
	Point() points.Point
	IsPoints() bool
	Points() points.Points
	IsLink() bool
	Link() links.Link
	IsLinks() bool
	Links() links.Links
}
