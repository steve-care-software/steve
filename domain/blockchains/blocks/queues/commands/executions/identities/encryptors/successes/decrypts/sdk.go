package decrypts

// Builder represents a decrypt builder
type Builder interface {
	Create() Builder
	WithVariable(variable string) Builder
	WithMessage(message []byte) Builder
	Now() (Decrypt, error)
}

// Decrypt represents a decrypt
type Decrypt interface {
	Variable() string
	Message() []byte
}
