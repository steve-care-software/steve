package publickeys

// PublicKey represents a public key
type PublicKey interface {
	Equals(pubKey PublicKey) bool
	Bytes() []byte
}
