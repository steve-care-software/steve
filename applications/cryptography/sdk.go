package cryptography

import "crypto/ed25519"

// NewApplication creates a new application
func NewApplication() Application {
	return createApplication()
}

// Application represents a cryptography application
type Application interface {
	Encrypt(message []byte, password []byte) ([]byte, error)       // encrypt data using a password
	Decrypt(cipher []byte, password []byte) ([]byte, error)        // decrypt a cipher using a password
	GeneratePrivateKey(words []string) (ed25519.PrivateKey, error) // generate a private key and returns it
}
