package paths

import "github.com/steve-care-software/steve/domain/connections"

type paths struct {
	list []Path
}

func createPaths(
	list []Path,
) Paths {
	out := paths{
		list: list,
	}

	return &out
}

// List returns the list of paths
func (obj *paths) List() []Path {
	return obj.list
}

// Successfuls returns the successful paths
func (obj *paths) Successfuls() [][]connections.Connection {
	output := [][]connections.Connection{}
	for _, onePath := range obj.list {
		successfuls := onePath.Successfuls()
		output = append(output, successfuls...)
	}

	return output
}
