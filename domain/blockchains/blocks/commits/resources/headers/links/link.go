package links

type link struct {
	container []string
	name      string
}

func createLink(
	container []string,
	name string,
) Link {
	out := link{
		container: container,
		name:      name,
	}

	return &out
}

// Container returns the container
func (obj *link) Container() []string {
	return obj.container
}

// Name returns the name
func (obj *link) Name() string {
	return obj.name
}
