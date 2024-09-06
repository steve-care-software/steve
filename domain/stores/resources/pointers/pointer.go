package pointers

type pointer struct {
	index  uint
	length uint
}

func createPointer(
	index uint,
	length uint,
) Pointer {
	out := pointer{
		index:  index,
		length: length,
	}

	return &out
}

// Index returns the index
func (obj *pointer) Index() uint {
	return obj.index
}

// Length returns the length
func (obj *pointer) Length() uint {
	return obj.length
}
