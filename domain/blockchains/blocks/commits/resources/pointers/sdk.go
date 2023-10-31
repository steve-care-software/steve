package pointers

// Builder represents a pointer builder
type Builder interface {
	Create() Builder
	WithIndex(index uint) Builder
	WithLength(length uint) Builder
	Now() (Pointer, error)
}

// Pointer represents a data pointer
type Pointer interface {
	Index() uint
	Length() uint
	Next() uint
	Delimiter() uint
}
