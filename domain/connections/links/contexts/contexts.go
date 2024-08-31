package contexts

type contexts struct {
	list []Context
}

func createContexts(
	list []Context,
) Contexts {
	out := contexts{
		list: list,
	}

	return &out
}

// List returns the list of contexts
func (obj *contexts) List() []Context {
	return obj.list
}
