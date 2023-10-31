package deletes

// Builder represents a delete builder
type Builder interface {
	Create() Builder
	WithIndex(index uint) Builder
	Now() (Delete, error)
}

// Delete represents a delete identity
type Delete interface {
	Index() uint
}
