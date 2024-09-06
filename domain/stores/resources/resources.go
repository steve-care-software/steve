package resources

type resources struct {
	list []Resource
}

func createResources(
	list []Resource,
) Resources {
	out := resources{
		list: list,
	}

	return &out
}

// List returns the list of resources
func (obj *resources) List() []Resource {
	return obj.list
}
