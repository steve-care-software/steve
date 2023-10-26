package passwords

// Application represents a password application
type Application interface {
	Encrypt(input []byte, password []byte) ([]byte, error)
	Decrypt(cipher []byte, password []byte) ([]byte, error)
}
