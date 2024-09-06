package pointers

// NewPointerForTests creates a new pointer for tests
func NewPointerForTests(index uint, length uint) Pointer {
	ins, err := NewBuilder().Create().
		WithIndex(index).
		WithLength(length).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
