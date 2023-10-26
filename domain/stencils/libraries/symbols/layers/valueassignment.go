package layers

type valueAssignment struct {
	name  string
	value Value
}

func createValueAssignment(
	name string,
	value Value,
) ValueAssignment {
	out := valueAssignment{
		name:  name,
		value: value,
	}

	return &out
}

// Name returns the name
func (obj *valueAssignment) Name() string {
	return obj.name
}

// Value returns the value
func (obj *valueAssignment) Value() Value {
	return obj.value
}
