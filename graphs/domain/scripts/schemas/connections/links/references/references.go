package references

type references struct {
	list []Reference
}

func createReferences(
	list []Reference,
) References {
	out := references{
		list: list,
	}

	return &out
}

// List returns the list of reference
func (obj *references) List() []Reference {
	return obj.list
}
