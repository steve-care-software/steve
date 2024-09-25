package parameters

// Parameters represents a func call parameters
type Parameters interface {
	List() []Parameter
}

// Parameter represents a func call parameter
type Parameter interface {
	Current() string
	Local() string
}
