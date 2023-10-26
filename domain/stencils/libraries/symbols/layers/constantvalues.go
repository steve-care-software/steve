package layers

type constantValues struct {
	list []ConstantValue
}

func createConstantValues(
	list []ConstantValue,
) ConstantValues {
	out := constantValues{
		list: list,
	}

	return &out
}

// List returns the constantValues
func (obj *constantValues) List() []ConstantValue {
	return obj.list
}
