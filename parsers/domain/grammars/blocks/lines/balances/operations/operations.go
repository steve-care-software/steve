package operations

type operations struct {
	list []Operation
}

func createOprations(
	list []Operation,
) Operations {
	out := operations{
		list: list,
	}

	return &out
}

// List returns the list of operations
func (obj *operations) List() []Operation {
	return obj.list
}
