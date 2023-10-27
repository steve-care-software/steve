package assignments

// Builder represents an assignment builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithValue(value []byte) Builder
	IsInternal() Builder
	Now() (Assignment, error)
}

// Assignment represents an assignment
type Assignment interface {
	Name() string
	Value() []byte
}

// Content represents an assignment content
type Content interface {
	IsInternal() bool
	IsValue() bool
	Value() []byte
}
