package permissions

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the permission builder
type Builder interface {
	Create() Builder
	WithNames(names []string) Builder
	WithCompensation(compensation float64) Builder
	Now() (Permission, error)
}

// Permission represents a permission
type Permission interface {
	Names() []string
	HasCompensation() bool
	Compensation() float64
}
