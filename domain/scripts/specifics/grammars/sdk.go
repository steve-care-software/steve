package grammars

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/programs/grammars/rules"
	"github.com/steve-care-software/steve/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/domain/scripts/specifics/grammars/blocks"
	"github.com/steve-care-software/steve/domain/scripts/specifics/grammars/constants"
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
	WithHead(head heads.Head) Builder
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
	Head() heads.Head
	Entry() string
	Rules() rules.Rules
	Blocks() blocks.Blocks
	HasOmit() bool
	Omit() []string
	HasConstants() bool
	Constants() constants.Constants
}
