package instructions

type constant struct {
	name  string
	value []byte
}

func createConstant(
	name string,
	value []byte,
) Constant {
	out := constant{
		name:  name,
		value: value,
	}

	return &out
}

// Name returns the name
func (obj *constant) Name() string {
	return obj.name
}

// Value returns the value
func (obj *constant) Value() []byte {
	return obj.value
}
