package votes

// Builder represents a vote builder
type Builder interface {
	Create() Builder
	WithMessage(message []byte) Builder
	Now() (Vote, error)
}

// Vote represents a vote
type Vote interface {
	Message() []byte
}
