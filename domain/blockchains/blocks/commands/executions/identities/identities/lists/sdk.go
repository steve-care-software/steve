package lists

// Builder represents a list builder
type Builder interface {
	Create() Builder
	WithVariable(variable string) Builder
	WithUsernames(usernames []string) Builder
	Now() (List, error)
}

// List represents a list
type List interface {
	Variable() string
	Usernames() []string
}
