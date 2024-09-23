package grammars

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/programs/grammars/rules"
	"github.com/steve-care-software/steve/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/domain/scripts/specifics/grammars/blocks"
	"github.com/steve-care-software/steve/domain/scripts/specifics/grammars/constants"
)

// Grammar represents a grammar
type Grammar interface {
	Hash() hash.Hash
	Head() heads.Head
	Entry() string
	Omit() []string
	Rules() rules.Rules
	Blocks() blocks.Blocks
	HasConstants() bool
	Constants() constants.Constants
}
