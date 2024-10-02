package headers

// Header represents the language header
type Header interface {
	Version() uint
	Name() string
}
