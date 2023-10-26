package layers

type valueAssignments struct {
	list []ValueAssignment
}

func createValueAssignments(
	list []ValueAssignment,
) ValueAssignments {
	out := valueAssignments{
		list: list,
	}

	return &out
}

// List returns the valueAssignments
func (obj *valueAssignments) List() []ValueAssignment {
	return obj.list
}
