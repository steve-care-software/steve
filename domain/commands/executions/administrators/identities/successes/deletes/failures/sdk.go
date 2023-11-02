package failures

// Builder represents a failure builder
type Builder interface {
	Create() Builder
	WithIndex(index uint) Builder
	IsIndexExceedAmount() Builder
	Now() (Failure, error)
}

// Failure represents a failure
type Failure interface {
	Index() uint
	Content() Content
}

// Content represents content
type Content interface {
	IsIndexExceedAmount() bool
}
