package containers

// Containers represents containers
type Containers interface {
	List() []Container
}

// Container represents a container
type Container interface {
	Flag() uint8 // single, vector, list, set, sorted_set
	Kind() Kind
}

// Kind represents a kind
type Kind interface {
	IsNumeric() bool
	Numeric() NumericKind
	IsEngine() bool
	Engine() *uint8 // path, routes, route
	IsRemaining() bool
	Remaining() *uint8 // string, bool
}

// NumericKind represents a numeric kind
type NumericKind interface {
	Flag() uint8 // uint, int, float
	Size() uint8 // 8, 16, 32, 64
}
