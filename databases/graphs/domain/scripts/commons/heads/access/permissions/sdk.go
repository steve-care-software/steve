package permissions

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewPermissionBuilder creates a new permission builder
func NewPermissionBuilder() PermissionBuilder {
	return createPermissionBuilder()
}

// Builder represents the permissions builder
type Builder interface {
	Create() Builder
	WithList(list []Permission) Builder
	Now() (Permissions, error)
}

// Permissions represents the permissions
type Permissions interface {
	List() []Permission
}

// PermissionBuilder represents the permission builder
type PermissionBuilder interface {
	Create() PermissionBuilder
	WithName(name string) PermissionBuilder
	WithCompensation(compensation float64) PermissionBuilder
	Now() (Permission, error)
}

// Permission represents a permission
type Permission interface {
	Name() string
	Compensation() float64
}
