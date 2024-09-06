package pointers

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewPointerBuilder creates a new pointer builder
func NewPointerBuilder() PointerBuilder {
	return createPointerBuilder()
}

// Adapter represents a pointer adapter
type Adapter interface {
	InstancesToBytes(ins Pointers) ([]byte, error)
	BytesToInstances(data []byte) (Pointers, error)
	InstanceToBytes(ins Pointer) ([]byte, error)
	BytesToInstance(data []byte) (Pointer, error)
}

// Builder represents the pointers builder
type Builder interface {
	Create() Builder
	WithList(list []Pointer) Builder
	Now() (Pointers, error)
}

// Pointers represents pointers
type Pointers interface {
	List() []Pointer
}

// PointerBuilder represents the pointer builder
type PointerBuilder interface {
	Create() PointerBuilder
	WithIndex(index uint) PointerBuilder
	WithLength(length uint) PointerBuilder
	Now() (Pointer, error)
}

// Pointer represents a pointer
type Pointer interface {
	Index() uint
	Length() uint
}
