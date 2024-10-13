package grammars

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/grammars/blocks"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/grammars/constants"
	"github.com/steve-care-software/steve/hash"
	"github.com/steve-care-software/steve/parsers/domain/grammars/rules"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the grammar builder
type Builder interface {
	Create() Builder
	WithEntry(entry string) Builder
	WithOmit(omit []string) Builder
	WithRules(rules rules.Rules) Builder
	WithBlocks(blocks blocks.Blocks) Builder
	WithConstants(constants constants.Constants) Builder
	Now() (Grammar, error)
}

// Grammar represents a grammar
type Grammar interface {
	Hash() hash.Hash
	Entry() string
	Rules() rules.Rules
	Blocks() blocks.Blocks
	HasOmit() bool
	Omit() []string
	HasConstants() bool
	Constants() constants.Constants
}
