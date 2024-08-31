package contexts

type context struct {
	name   string
	parent Context
}

func createContext(
	name string,
) Context {
	return createContextInternally(name, nil)
}

func createContextWithParent(
	name string,
	parent Context,
) Context {
	return createContextInternally(name, parent)
}

func createContextInternally(
	name string,
	parent Context,
) Context {
	out := context{
		name:   name,
		parent: parent,
	}

	return &out
}

// Name returns the name
func (obj *context) Name() string {
	return obj.name
}

// HasParent returns true if there is parent, false otherwise
func (obj *context) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *context) Parent() Context {
	return obj.parent
}
