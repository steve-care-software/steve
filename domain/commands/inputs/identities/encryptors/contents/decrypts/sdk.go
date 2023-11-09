package decrypts

// Builder represents a decrypt builder
type Builder interface {
	Create() Builder
	WithCipher(cipher []byte) Builder
	Now() (Decrypt, error)
}

// Decrypt represents a decrypt
type Decrypt interface {
	Cipher() []byte
}
