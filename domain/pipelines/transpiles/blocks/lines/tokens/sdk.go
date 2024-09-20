package tokens

import (
	"github.com/steve-care-software/steve/domain/pipelines/transpiles/blocks/lines/tokens/pointers"
	"github.com/steve-care-software/steve/domain/pipelines/transpiles/blocks/lines/tokens/updates"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewTokenBuilder creates a new token builder
func NewTokenBuilder() TokenBuilder {
	return createTokenBuilder()
}

// Builder represents tokens builder
type Builder interface {
	Create() Builder
	WithList(list []Token) Builder
	Now() (Tokens, error)
}

// Tokens represents tokens
type Tokens interface {
	List() []Token
}

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithUpdate(update updates.Update) TokenBuilder
	WithInsert(insert pointers.Pointer) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	IsUpdate() bool
	Update() updates.Update
	IsInsert() bool
	Insert() pointers.Pointer
}
