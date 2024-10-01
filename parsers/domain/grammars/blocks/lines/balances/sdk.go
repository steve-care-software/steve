package balances

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/operations"
)

// Builder represents the balance builder
type Builder interface {
	Create() Builder
	WithLines(lines []operations.Operations) Builder
	Now() (Balance, error)
}

// Balance represents balance
type Balance interface {
	Hash() hash.Hash
	Lines() []operations.Operations
}
