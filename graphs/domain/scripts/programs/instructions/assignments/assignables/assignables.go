package assignables

type assignables struct {
	list []Assignable
}

func createAssignables(
	list []Assignable,
) Assignables {
	out := assignables{
		list: list,
	}

	return &out
}

// List returns the list of assignable
func (obj *assignables) List() []Assignable {
	return obj.list
}
