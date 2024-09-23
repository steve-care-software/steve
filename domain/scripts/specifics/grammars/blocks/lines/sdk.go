package lines

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/grammars/blocks/lines/tokens"
)

// Lines represents lines
type Lines interface {
	Hash() hash.Hash
	List() []tokens.Tokens
}
