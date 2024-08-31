package paths

import "github.com/steve-care-software/steve/domain/connections"

type path struct {
	possibilities Paths
	destination   connections.Connection
}

func createPath(
	possibilities Paths,
) Path {
	return createPathInternally(possibilities, nil)
}

func createPathInternally(
	possibilities Paths,
	destination connections.Connection,
) Path {
	out := path{
		possibilities: possibilities,
		destination:   destination,
	}

	return &out
}

// Successfuls returns the successful paths
func (obj *path) Successfuls() [][]connections.Connection {
	if !obj.HasDestination() {
		return [][]connections.Connection{}
	}

	output := [][]connections.Connection{}
	list := obj.possibilities.List()
	for _, onePaths := range list {
		successfuls := onePaths.Successfuls()
		output = append(output, successfuls...)
	}

	return output
}

// Possibilities returns the possibilities
func (obj *path) Possibilities() Paths {
	return obj.possibilities
}

// HasDestination returns true if there is a destination, false otherwise
func (obj *path) HasDestination() bool {
	return obj.destination != nil
}

// Destination returns the destination, if any
func (obj *path) Destination() connections.Connection {
	return obj.destination
}
