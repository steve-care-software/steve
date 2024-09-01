package connections

import "strings"

type connections struct {
	list []Connection
}

func createConnections(
	list []Connection,
) Connections {
	out := connections{
		list: list,
	}

	return &out
}

// List returns the list of connections
func (obj *connections) List() []Connection {
	return obj.list
}

// Debug returns the string debug representation of the link
func (obj *connections) Debug() string {
	output := []string{}
	for _, oneConnection := range obj.list {
		output = append(output, oneConnection.Debug())
	}

	return strings.Join(output, "")
}
