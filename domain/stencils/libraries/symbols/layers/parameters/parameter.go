package parameters

type parameter struct {
	name string
	kind uint8
}

func createParameter(
	name string,
	kind uint8,
) Parameter {
	out := parameter{
		name: name,
		kind: kind,
	}

	return &out
}

// Name returns the name
func (obj *parameter) Name() string {
	return obj.name
}

// Kind returns the kind
func (obj *parameter) Kind() uint8 {
	return obj.kind
}
