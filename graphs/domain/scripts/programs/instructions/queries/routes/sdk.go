package routes

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links"
)

// Route represents a route
type Route interface {
	IsOptimal() bool
	Link() links.Link
}
