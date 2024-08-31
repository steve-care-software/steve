package routes

import "github.com/steve-care-software/steve/domain/connections"

type route struct {
	possibilities []connections.Connections
}

func createRoute(
	possibilities []connections.Connections,
) Route {
	out := route{
		possibilities: possibilities,
	}

	return &out
}

// Possibilities returns the possibilities
func (obj *route) Possibilities() []connections.Connections {
	return obj.possibilities
}
