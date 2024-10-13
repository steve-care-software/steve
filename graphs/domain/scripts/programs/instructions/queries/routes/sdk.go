package routes

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the route builder
type Builder interface {
	Create() Builder
	WithLink(link links.Link) Builder
	Now() (Route, error)
}

// Route represents a route
type Route interface {
	IsOptimal() bool
	Link() links.Link
}
