package permissions

// Permissions represents permissions
type Permissions interface {
	List() []Permission
}

// Permission represents a permission
type Permission interface {
	Path() []string
	CanRead() bool
	CanWrite() bool
	CanExecute() bool
}
