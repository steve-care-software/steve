package kinds

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a kind builder
type Builder interface {
	Create() Builder
	WithExecute(execute []string) Builder
	IsContinue() Builder
	IsPrompt() Builder
	Now() (Kind, error)
}

// Kind represents the return kind
type Kind interface {
	IsContinue() bool
	IsPrompt() bool
	IsExecute() bool
	Execute() []string
}
