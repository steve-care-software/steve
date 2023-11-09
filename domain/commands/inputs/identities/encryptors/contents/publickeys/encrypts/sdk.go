package encrypts

// Builder represents an encrypt builder
type Builder interface {
	Create() Builder
	WithMessage(message []byte) Builder
	Now() (Encrypt, error)
}

// Encrypt represents an encrypt
type Encrypt interface {
	Message() []byte
}
