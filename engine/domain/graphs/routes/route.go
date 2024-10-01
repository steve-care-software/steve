package routes

import (
	"fmt"
	"strings"

	"github.com/steve-care-software/steve/engine/domain/graphs/connections"
)

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

// Debug returns the debug string representation of the route
func (obj *route) Debug() string {
	output := []string{}
	for idx, onePossibility := range obj.possibilities {
		str := fmt.Sprintf("\n++++++      %d      ++++++\n%s\n++++++", idx, onePossibility.Debug())
		output = append(output, str)
	}

	return strings.Join(output, "")
}
