package targets

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the target builder
type Builder interface {
	Create() Builder
	WithConstant(constant string) Builder
	WithRule(rule string) Builder
	Now() (Target, error)
}

// Target represents a target
type Target interface {
	IsConstant() bool
	Constant() string
	IsRule() bool
	Rule() string
}
