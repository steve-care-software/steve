package identities

// Retrieve represents a retrieve
type Retrieve interface {
	User() string
	Password() []byte
}

// Insert represents an insert
type Insert interface {
	User() string
	Password() []byte
	Name() string
	Description() string
}

// Update represents an update
type Update interface {
	User() string
	CurrentPassword() []byte
	HasNewPassword() bool
	NewPassword() []byte
	HasName() bool
	Name() string
	HasDescription() bool
	Description() string
}

// Delete represents a delete
type Delete interface {
	User() string
	Password() []byte
}
