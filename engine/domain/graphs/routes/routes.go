package routes

type routes struct {
	list []Route
}

func createRoutes(
	list []Route,
) Routes {
	out := routes{
		list: list,
	}

	return &out
}

// List returns the list of routes
func (obj *routes) List() []Route {
	return obj.list
}
