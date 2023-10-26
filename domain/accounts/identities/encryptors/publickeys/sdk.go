package publickeys

// PublicKey represents a public key
type PublicKey interface {
	Encrypt(msg []byte) ([]byte, error)
	Bytes() []byte
}
