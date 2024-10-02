package routes

import "github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/links/references"

// Routes represents routes
type Routes interface {
	List() []Route
}

// Route represents a route
type Route interface {
	Possibilities() []references.References
}
