package routes

import "github.com/steve-care-software/steve/engine/domain/graphs/connections"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewRouteBuilder creates a new route builder
func NewRouteBuilder() RouteBuilder {
	return createRouteBuilder()
}

// Builder represents routes builder
type Builder interface {
	Create() Builder
	WithList(list []Route) Builder
	Now() (Routes, error)
}

// Routes represents routes
type Routes interface {
	List() []Route
}

// RouteBuilder represents a route builder
type RouteBuilder interface {
	Create() RouteBuilder
	WithPossibilities(possibilities []connections.Connections) RouteBuilder
	Now() (Route, error)
}

// Route represents a route
type Route interface {
	Possibilities() []connections.Connections
	Debug() string
}
