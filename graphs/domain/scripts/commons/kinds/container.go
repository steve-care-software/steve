package kinds

// container represents the implementation of the Container interface
type container struct {
	flag uint8
	kind Kind
}

func createContainer(flag uint8, kind Kind) Container {
	return &container{
		flag: flag,
		kind: kind,
	}
}

// Flag returns the flag of the container
func (obj *container) Flag() uint8 {
	return obj.flag
}

// Kind returns the kind of the container
func (obj *container) Kind() Kind {
	return obj.kind
}
