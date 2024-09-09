package resources

const appNotInitErr = "the application has not been initialized yet"
const alreadyModifiedErrPattern = "the resource (identifier: %s) has already been modified in the current session"

// Application represents the resources application
type Application interface {
	Init(dbIdentifier string) error
	Retrieve(identifier string) ([]byte, error)
	Insert(identifier string, data []byte) error
	Save(identifier string, data []byte) error
	Delete(identifier string) error
	Commit() error
	Cancel() error
	Rollback(amount uint) error
}
