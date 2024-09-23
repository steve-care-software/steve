package tokens

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/grammars/constants/tokens/elements"
)

// Tokens represents tokens
type Tokens interface {
	Hash() hash.Hash
	List() []Token
}

// Token represents a token
type Token interface {
	Hash() hash.Hash
	Element() elements.Element
	Occurences() uint
}
