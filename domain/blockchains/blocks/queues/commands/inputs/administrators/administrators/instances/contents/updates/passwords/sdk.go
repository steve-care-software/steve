package passwords

// Builder represents a password builder
type Builder interface {
	Create() Builder
	WithCurrent(current []byte) Builder
	WithUpdated(updated []byte) Builder
	Now() (Password, error)
}

// Password represents a password
type Password interface {
	Current() []byte
	Updated() []byte
}
