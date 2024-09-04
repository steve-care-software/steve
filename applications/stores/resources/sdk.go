package resources

// Application represents the resources application
type Application interface {
	Retrieve(identifier string) ([]byte, error)
	Insert(identifier string, data []byte) error
	Save(identifier string, data []byte) error
	Delete(identifier string) error
}
