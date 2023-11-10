package properties

// Builder represents the property builder
type Builder interface {
	Create() Builder
	IsDashboard() Builder
	IsEncryptor() Builder
	IsSigner() Builder
	IsProfile() Builder
	IsHasConnections() Builder
	IsConnections() Builder
	IsHasShares() Builder
	IsShares() Builder
	Now() (Property, error)
}

// Property represents a property
type Property interface {
	IsDashboard() bool
	IsEncryptor() bool
	IsSigner() bool
	IsProfile() bool
	IsHasConnections() bool
	IsConnections() bool
	IsHasShares() bool
	IsShares() bool
}
