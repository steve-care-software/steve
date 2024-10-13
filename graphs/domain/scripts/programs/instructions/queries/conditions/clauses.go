package conditions

type clauses struct {
	list []Clause
}

func createClauses(
	list []Clause,
) Clauses {
	out := clauses{
		list: list,
	}

	return &out
}

// List returns the list of clause
func (obj *clauses) List() []Clause {
	return obj.list
}
