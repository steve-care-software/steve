package messages

// Message represents a message
type Message interface {
	Bytes() []byte
	HasSignature() bool
	Signature() []byte
}
