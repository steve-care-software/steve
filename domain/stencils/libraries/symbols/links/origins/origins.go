package origins

type origins struct {
	list []Origin
}

func createOrigins(
	list []Origin,
) Origins {
	out := origins{
		list: list,
	}

	return &out
}

// List returns the origins
func (obj *origins) List() []Origin {
	return obj.list
}
