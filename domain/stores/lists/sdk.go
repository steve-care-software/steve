package lists

// Adapter represents a list adapter
type Adapter interface {
	ToBytes(list [][]byte) ([]byte, error)
	ToInstance(data []byte) ([][]byte, error)
}
