package messages

// Builder represents a message builder
type Builder interface {
	Create() Builder
	WithBytes(bytes []byte) Builder
	WithSignature(sig []byte) Builder
	Now() (Message, error)
}

// Message represents a message
type Message interface {
	Bytes() []byte
	HasSignature() bool
	Signature() []byte
}
