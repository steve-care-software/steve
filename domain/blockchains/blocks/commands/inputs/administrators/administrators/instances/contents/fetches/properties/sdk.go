package properties

// Builder represents a property builder
type Builder interface {
	Create() Builder
	IsUsername() Builder
	IsDashboard() Builder
	IsHasIdentities() Builder
	IsIdentities() Builder
	Now() (Property, error)
}

// Property represents a property
type Property interface {
	IsUsername() bool
	IsDashboard() bool
	IsHasIdentities() bool
	IsIdentities() bool
}
