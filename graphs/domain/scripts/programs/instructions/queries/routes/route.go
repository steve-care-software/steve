package routes

import "github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links"

type route struct {
	link      links.Link
	isOptimal bool
}

func createRoute(link links.Link, isOptimal bool) Route {
	return &route{
		link:      link,
		isOptimal: isOptimal,
	}
}

// IsOptimal returns true if the route is optimal
func (obj *route) IsOptimal() bool {
	return obj.isOptimal
}

// Link returns the link of the route
func (obj *route) Link() links.Link {
	return obj.link
}
