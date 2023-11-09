package deletes

// Builder represents a delete builder
type Builder interface {
	Create() Builder
	WithPassword(password []byte) Builder
	Now() (Delete, error)
}

// Delete represents a delete
type Delete interface {
	Password() []byte
}
