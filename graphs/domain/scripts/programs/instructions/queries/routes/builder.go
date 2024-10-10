package routes

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links"
)

type builder struct {
	link      links.Link
	isOptimal bool
}

func createBuilder() Builder {
	return &builder{
		link:      nil,
		isOptimal: false,
	}
}

// Create initializes the route builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithLink adds a link to the route builder
func (app *builder) WithLink(link links.Link) Builder {
	app.link = link
	return app
}

// Now builds a new Route instance
func (app *builder) Now() (Route, error) {
	if app.link == nil {
		return nil, errors.New("the link is mandatory in order to build a Route instance")
	}

	return createRoute(app.link, app.isOptimal), nil
}
