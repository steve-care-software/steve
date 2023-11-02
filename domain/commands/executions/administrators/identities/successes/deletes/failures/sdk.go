package failures

// Builder represents a failure builder
type Builder interface {
	Create() Builder
	WithIndex(index uint) Builder
	WithAmount(amount uint) Builder
	IsIndexExceedAmount() Builder
	Now() (Failure, error)
}

// Failure represents a failure
type Failure interface {
	Index() uint
	Amount() uint
	Content() Content
}

// Content represents content
type Content interface {
	IsIndexExceedAmount() bool
}
