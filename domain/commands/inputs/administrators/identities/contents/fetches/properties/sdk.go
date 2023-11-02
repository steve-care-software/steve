package properties

// Builder represents a property builder
type Builder interface {
	Create() Builder
	WithAtIndex(atIndex uint) Builder
	IsList() Builder
	IsAmount() Builder
	Now() (Property, error)
}

// Property represents a property
type Property interface {
	IsList() bool
	IsAmount() bool
	IsAtIndex() bool
	AtIndex() *uint
}
