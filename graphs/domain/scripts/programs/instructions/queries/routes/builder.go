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
func (obj *builder) Create() Builder {
	return createBuilder()
}

// WithLink adds a link to the route builder
func (obj *builder) WithLink(link links.Link) Builder {
	obj.link = link
	return obj
}

// Now builds a new Route instance
func (obj *builder) Now() (Route, error) {
	if obj.link == nil {
		return nil, errors.New("the link is mandatory in order to build a Route instance")
	}

	return createRoute(obj.link, obj.isOptimal), nil
}
