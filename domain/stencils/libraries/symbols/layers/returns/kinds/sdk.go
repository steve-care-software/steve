package kinds

import "github.com/steve-care-software/steve/domain/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
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
	Hash() hash.Hash
	IsContinue() bool
	IsPrompt() bool
	IsExecute() bool
	Execute() []string
}
