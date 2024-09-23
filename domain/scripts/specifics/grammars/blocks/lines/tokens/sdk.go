package tokens

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/programs/grammars/blocks/lines/tokens/cardinalities"
	"github.com/steve-care-software/steve/domain/programs/grammars/blocks/lines/tokens/reverses"
	"github.com/steve-care-software/steve/domain/scripts/specifics/grammars/blocks/lines/tokens/elements"
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
	Cardinality() cardinalities.Cardinality
	HasReverse() bool
	Reverse() reverses.Reverse
}
