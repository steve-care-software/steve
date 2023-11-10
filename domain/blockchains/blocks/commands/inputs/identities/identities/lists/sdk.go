package lists

// Builder represents a list builder
type Builder interface {
	Create() Builder
	WithAssignToVariable(assignToVariable string) Builder
	Now() (List, error)
}

// List represents a list identity
type List interface {
	AssignToVariable() string
}
