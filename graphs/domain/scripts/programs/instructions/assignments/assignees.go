package assignments

type assignees struct {
	list []Assignee
}

func createAssignees(
	list []Assignee,
) Assignees {
	out := assignees{
		list: list,
	}

	return &out
}

// List returns the list of assignee
func (obj *assignees) List() []Assignee {
	return obj.list
}
