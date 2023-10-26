package credentials

// Credentials represents credentials
type Credentials interface {
	Username() string
	Password() []byte
}
