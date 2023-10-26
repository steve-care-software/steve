package connections

// Connection represents a connection
type Connection interface {
	Input() []byte
	Signature() []byte
}
