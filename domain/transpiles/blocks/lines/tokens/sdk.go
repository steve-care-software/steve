package tokens

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens/pointers"
	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens/updates"
)

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
	WithDelete(delete pointers.Pointer) TokenBuilder
	WithInsert(insert pointers.Pointer) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Hash() hash.Hash
	IsUpdate() bool
	Update() updates.Update
	IsDelete() bool
	Delete() pointers.Pointer
	IsInsert() bool
	Insert() pointers.Pointer
}
