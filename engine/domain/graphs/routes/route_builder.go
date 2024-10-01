package routes

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/graphs/connections"
)

type routeBuilder struct {
	possibilities []connections.Connections
}

func createRouteBuilder() RouteBuilder {
	out := routeBuilder{
		possibilities: nil,
	}

	return &out
}

// Create initializes the builder
func (app *routeBuilder) Create() RouteBuilder {
	return createRouteBuilder()
}

// WithPossibilities add possibilities to the builder
func (app *routeBuilder) WithPossibilities(possibilities []connections.Connections) RouteBuilder {
	app.possibilities = possibilities
	return app
}

// Now builds a new Route instance
func (app *routeBuilder) Now() (Route, error) {
	if app.possibilities != nil && len(app.possibilities) <= 0 {
		app.possibilities = nil
	}

	if app.possibilities == nil {
		return nil, errors.New("there must be at least 1 possible connections serie in order to build a Route instance")
	}

	return createRoute(
		app.possibilities,
	), nil
}
