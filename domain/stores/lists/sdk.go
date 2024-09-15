package lists

const remainingTooSmallPatternErr = "the remaining was expected to contain at least %d bytes in order to convert it to a list"

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	return createAdapter()
}

// Adapter represents a list adapter
type Adapter interface {
	ToBytes(list [][]byte) ([]byte, error)
	ToInstance(data []byte) ([][]byte, error)
}
