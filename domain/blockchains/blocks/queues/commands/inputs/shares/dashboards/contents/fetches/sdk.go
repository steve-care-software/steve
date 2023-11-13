package fetches

// Builder represents the fetch builder
type Builder interface {
	Create() Builder
	WithAssignTo(assignTo string) Builder
	IsRoot() Builder
	IsStencils() Builder
	Now() (Fetch, error)
}

// Fetch represents a fetch
type Fetch interface {
	AssignTo() string
	Content() Content
}

// Content represents a fetch content
type Content interface {
	IsRoot() bool
	IsStencils() bool
}
