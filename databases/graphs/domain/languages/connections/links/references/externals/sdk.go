package externals

// External represents an external reference
type External interface {
	Schema() string
	Point() string
}
