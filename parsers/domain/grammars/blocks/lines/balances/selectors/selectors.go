package selectors

type selectors struct {
	list []Selector
}

func createSelectors(
	list []Selector,
) Selectors {
	out := selectors{
		list: list,
	}

	return &out
}

// List returns the list of selectors
func (obj *selectors) List() []Selector {
	return obj.list
}
