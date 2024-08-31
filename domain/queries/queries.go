package queries

type queries struct {
	list []Query
}

func createQueries(
	list []Query,
) Queries {
	out := queries{
		list: list,
	}

	return &out
}

// List returns the list of queries
func (obj *queries) List() []Query {
	return obj.list
}
