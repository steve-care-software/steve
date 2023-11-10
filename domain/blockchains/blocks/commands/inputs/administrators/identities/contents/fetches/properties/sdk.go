package properties

// Builder represents a property builder
type Builder interface {
	Create() Builder
	WithAtIndex(atIndex uint) Builder
	IsAmount() Builder
	Now() (Property, error)
}

// Property represents a property
type Property interface {
	IsAmount() bool
	IsAtIndex() bool
	AtIndex() *uint
}
