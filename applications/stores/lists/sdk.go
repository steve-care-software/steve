package lists

// Application represents the list application
type Application interface {
	Amount(name string) (*uint, error)
	Retrieve(name string, index uint, amount uint) ([][]byte, error)
	Append(name string, values [][]byte) error
	InsertAtIndex(name string, index uint, values [][]byte) error
	Remove(name string, index uint, amount uint) error
}
