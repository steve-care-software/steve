package credentials

// Builder represents a credentials builder
type Builder interface {
	Create() Builder
	WithUsername(username string) Builder
	WithPassword(password []byte) Builder
	Now() (Credentials, error)
}

// Credentials represents credentials
type Credentials interface {
	Username() string
	Password() []byte
}
