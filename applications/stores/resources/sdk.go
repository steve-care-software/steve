package resources

const appNotInitErr = "the application has not been initialized yet"

// Application represents the resources application
type Application interface {
	Init(dbIdentifier string) error
	Retrieve(identifier string) ([]byte, error)
	Insert(identifier string, data []byte) error
	Save(identifier string, data []byte)
	Delete(identifier string) error
	Commit() error
	Cancel() error
	Rollback(amount uint) error
}
