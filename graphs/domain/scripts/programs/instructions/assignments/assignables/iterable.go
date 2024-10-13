package assignables

type iterable struct {
	listMap  ListMap
	variable string
}

func createIterableWithListMap(listMap ListMap) Iterable {
	return createIterableInternally(listMap, "")
}

func createIterableWithVariable(variable string) Iterable {
	return createIterableInternally(nil, variable)
}

func createIterableInternally(listMap ListMap, variable string) Iterable {
	return &iterable{
		listMap:  listMap,
		variable: variable,
	}
}

// IsListMap returns true if the iterable contains a ListMap
func (obj *iterable) IsListMap() bool {
	return obj.listMap != nil
}

// ListMap returns the ListMap if present
func (obj *iterable) ListMap() ListMap {
	return obj.listMap
}

// IsVariable returns true if the iterable contains a variable
func (obj *iterable) IsVariable() bool {
	return obj.variable != ""
}

// Variable returns the variable if present
func (obj *iterable) Variable() string {
	return obj.variable
}
