package signs

// Builder represents a sign builder
type Builder interface {
	Create() Builder
	WithMessage(message []byte) Builder
	Now() (Sign, error)
}

// Sign represents a sign
type Sign interface {
	Message() []byte
}
