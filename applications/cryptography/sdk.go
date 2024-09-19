package cryptography

import "crypto/ed25519"

const (
	// LangEnglish represents the english language
	LangEnglish (uint8) = iota

	// LangFrench represents the french language
	LangFrench
)

// NewApplication creates a new application
func NewApplication() Application {
	return createApplication()
}

// Application represents a cryptography application
type Application interface {
	Encrypt(message []byte, password []byte) ([]byte, error)                 // encrypt data using a password
	Decrypt(cipher []byte, password []byte) ([]byte, error)                  // decrypt a cipher using a password
	GeneratePrivateKey(language uint8) (ed25519.PrivateKey, []string, error) // generate a private key and returns it
	GeneratePrivateKeyFromSeedWords(seedWords []string) (ed25519.PrivateKey, error)
}
