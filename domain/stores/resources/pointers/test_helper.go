package pointers

// NewPointersForTests creates a new pointers for tests
func NewPointersForTests(list []Pointer) Pointers {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewPointerForTests creates a new pointer for tests
func NewPointerForTests(index uint, length uint) Pointer {
	ins, err := NewPointerBuilder().Create().
		WithIndex(index).
		WithLength(length).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
