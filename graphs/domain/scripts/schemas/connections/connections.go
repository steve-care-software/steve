package connections

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

// List returns the list of connection
func (obj *connections) List() []Connection {
	return obj.list
}
