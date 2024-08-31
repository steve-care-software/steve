package paths

import "github.com/steve-care-software/steve/domain/connections"

type paths struct {
	list []connections.Connections
}

func createPaths(
	list []connections.Connections,
) Paths {
	out := paths{
		list: list,
	}

	return &out
}

// List returns the list of path
func (obj *paths) List() []connections.Connections {
	return obj.list
}
