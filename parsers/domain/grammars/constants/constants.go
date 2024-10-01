package constants

type constants struct {
	list []Constant
}

func createOprations(
	list []Constant,
) Constants {
	out := constants{
		list: list,
	}

	return &out
}

// List returns the list of constants
func (obj *constants) List() []Constant {
	return obj.list
}
