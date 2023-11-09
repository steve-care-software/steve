package encrypts

// Builder represents an encrypt builder
type Builder interface {
	Create() Builder
	WithCipher(cipher []byte) Builder
	Now() (Encrypt, error)
}

// Encrypt represents an encrypt
type Encrypt interface {
	Cipher() []byte
}
