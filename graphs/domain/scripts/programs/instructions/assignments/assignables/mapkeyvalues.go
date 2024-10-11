package assignables

type mapKeyValues struct {
	list []MapKeyValue
}

func createMapKeyValues(
	list []MapKeyValue,
) MapKeyValues {
	out := mapKeyValues{
		list: list,
	}

	return &out
}

// List returns the list of mapKeyValue
func (obj *mapKeyValues) List() []MapKeyValue {
	return obj.list
}
