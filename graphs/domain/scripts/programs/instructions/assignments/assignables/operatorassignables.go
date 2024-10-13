package assignables

type operatorAssignables struct {
	list []OperatorAssignable
}

func createOperatorAssignables(
	list []OperatorAssignable,
) OperatorAssignables {
	out := operatorAssignables{
		list: list,
	}

	return &out
}

// List returns the list of operatorAssignable
func (obj *operatorAssignables) List() []OperatorAssignable {
	return obj.list
}
