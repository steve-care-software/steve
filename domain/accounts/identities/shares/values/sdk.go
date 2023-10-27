package values

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/signers/votes"
	"github.com/steve-care-software/steve/domain/pointers"
	"github.com/steve-care-software/steve/domain/pointers/symbols"
)

// Values represents values
type Values interface {
	List() []Value
}

// Value represents the value of a share
type Value interface {
	Content() Content
	Vote() votes.Vote
}

// Content represents the value content
type Content interface {
	Pointer() pointers.Pointer
	Symbol() symbols.Symbol
}
