package kinds

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
)

// Builder represents the kind builder
type Builder interface {
	Create() Builder
	WithExecute(exec []string) Builder
	IsContinue() bool
	IsPrompt() bool
	Now() (Kind, error)
}

// Kind represents the kind
type Kind interface {
	Hash() hash.Hash
	IsContinue() bool
	IsPrompt() bool
	IsExecute() bool
	Execute() []string
}
