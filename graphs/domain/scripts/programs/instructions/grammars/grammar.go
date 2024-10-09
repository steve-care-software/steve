package grammars

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/grammars/blocks"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/grammars/constants"
	"github.com/steve-care-software/steve/hash"
	"github.com/steve-care-software/steve/parsers/domain/grammars/rules"
)

type grammar struct {
	hash      hash.Hash
	entry     string
	rules     rules.Rules
	blocks    blocks.Blocks
	omit      []string
	constants constants.Constants
}

func createGrammar(
	hash hash.Hash,
	entry string,
	rules rules.Rules,
	blocks blocks.Blocks,
) Grammar {
	return createGrammarInternally(
		hash,
		entry,
		rules,
		blocks,
		nil,
		nil,
	)
}

func createGrammarWithOmit(
	hash hash.Hash,
	entry string,
	rules rules.Rules,
	blocks blocks.Blocks,
	omit []string,
) Grammar {
	return createGrammarInternally(
		hash,
		entry,
		rules,
		blocks,
		omit,
		nil,
	)
}

func createGrammarWithConstants(
	hash hash.Hash,
	entry string,
	rules rules.Rules,
	blocks blocks.Blocks,
	constants constants.Constants,
) Grammar {
	return createGrammarInternally(
		hash,
		entry,
		rules,
		blocks,
		nil,
		constants,
	)
}

func createGrammarWithOmitAndConstants(
	hash hash.Hash,
	entry string,
	rules rules.Rules,
	blocks blocks.Blocks,
	omit []string,
	constants constants.Constants,
) Grammar {
	return createGrammarInternally(
		hash,
		entry,
		rules,
		blocks,
		omit,
		constants,
	)
}

func createGrammarInternally(
	hash hash.Hash,
	entry string,
	rules rules.Rules,
	blocks blocks.Blocks,
	omit []string,
	constants constants.Constants,
) Grammar {
	out := grammar{
		hash:      hash,
		entry:     entry,
		rules:     rules,
		blocks:    blocks,
		omit:      omit,
		constants: constants,
	}

	return &out
}

// Hash returns the hash
func (obj *grammar) Hash() hash.Hash {
	return obj.hash
}

// Entry returns the entry
func (obj *grammar) Entry() string {
	return obj.entry
}

// Rules returns the rules
func (obj *grammar) Rules() rules.Rules {
	return obj.rules
}

// Blocks returns the blocks
func (obj *grammar) Blocks() blocks.Blocks {
	return obj.blocks
}

// HasOmit returns true if there is omit, false otherwise
func (obj *grammar) HasOmit() bool {
	return obj.omit != nil
}

// Omit returns the omit, if any
func (obj *grammar) Omit() []string {
	return obj.omit
}

// HasConstants returns true if there is constants, false otherwise
func (obj *grammar) HasConstants() bool {
	return obj.constants != nil
}

// Constants returns the constants, if any
func (obj *grammar) Constants() constants.Constants {
	return obj.constants
}
