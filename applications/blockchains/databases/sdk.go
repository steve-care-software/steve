package databases

// OnOpenFn represents the onOpen func
type OnOpenFn func(context uint) error

// Application represents the database application
type Application interface {
	Exists(name string) (bool, error)
	New(name string) error
	Delete(name string) error
	Open(name string) (*uint, error)
	Lock(context uint) error
	Unlock(context uint) error
	Read(context uint, offset uint, length uint) ([]byte, error)
	Write(context uint, offset int64, data []byte) error
	Copy(context uint, destination string) error
	Close(context uint) error
}
