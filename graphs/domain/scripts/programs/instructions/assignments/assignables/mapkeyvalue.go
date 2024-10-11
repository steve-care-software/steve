package assignables

type mapKeyValue struct {
	name       string
	assignable Assignable
}

func createMapKeyValue(
	name string,
	assignable Assignable,
) MapKeyValue {
	out := mapKeyValue{
		name:       name,
		assignable: assignable,
	}

	return &out
}

// Name returns the name
func (obj *mapKeyValue) Name() string {
	return obj.name
}

// Assignable returns the assignable
func (obj *mapKeyValue) Assignable() Assignable {
	return obj.assignable
}
