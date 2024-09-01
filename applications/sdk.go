package applications

// Application represents the grammar application
type Application interface {
	Execute(input []byte) ([]byte, error)
}
