package assignables

type listMap struct {
	list Assignables
	mp   MapKeyValues
}

func createListMapWithList(list Assignables) ListMap {
	return createListMapInternally(list, nil)
}

func createListMapWithMap(mp MapKeyValues) ListMap {
	return createListMapInternally(nil, mp)
}

func createListMapInternally(list Assignables, mp MapKeyValues) ListMap {
	return &listMap{
		list: list,
		mp:   mp,
	}
}

// IsList returns true if the ListMap contains a list
func (obj *listMap) IsList() bool {
	return obj.list != nil
}

// List returns the list if present
func (obj *listMap) List() Assignables {
	return obj.list
}

// IsMap returns true if the ListMap contains a map
func (obj *listMap) IsMap() bool {
	return obj.mp != nil
}

// Map returns the map if present
func (obj *listMap) Map() MapKeyValues {
	return obj.mp
}
