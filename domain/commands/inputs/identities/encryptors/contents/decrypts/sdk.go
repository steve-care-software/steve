package decrypts

// Builder represents a decrypt builder
type Builder interface {
	Create() Builder
	WithAssignToVariable(assignToVariable string) Builder
	WithCipher(cipher []byte) Builder
	Now() (Decrypt, error)
}

// Decrypt represents a decrypt
type Decrypt interface {
	AssignToVariable() string
	Cipher() []byte
}
