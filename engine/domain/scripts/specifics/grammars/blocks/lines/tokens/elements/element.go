package elements

import (
	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/grammars/blocks/lines/tokens/elements/references"
)

type element struct {
	hash      hash.Hash
	reference references.Reference
	rule      string
	constant  string
	block     string
}

func createElementWithReference(
	hash hash.Hash,
	reference references.Reference,
) Element {
	return createElementInternally(
		hash,
		reference,
		"",
		"",
		"",
	)
}

func createElementWithRule(
	hash hash.Hash,
	rule string,
) Element {
	return createElementInternally(
		hash,
		nil,
		rule,
		"",
		"",
	)
}

func createElementWithConstant(
	hash hash.Hash,
	constant string,
) Element {
	return createElementInternally(
		hash,
		nil,
		"",
		constant,
		"",
	)
}

func createElementWithBlock(
	hash hash.Hash,
	block string,
) Element {
	return createElementInternally(
		hash,
		nil,
		"",
		"",
		block,
	)
}

func createElementInternally(
	hash hash.Hash,
	reference references.Reference,
	rule string,
	constant string,
	block string,
) Element {
	out := element{
		hash:      hash,
		reference: reference,
		rule:      rule,
		constant:  constant,
		block:     block,
	}

	return &out
}

// Hash returns the hash
func (obj *element) Hash() hash.Hash {
	return obj.hash
}

// IsReference returns true if there is a reference, false otherwise
func (obj *element) IsReference() bool {
	return obj.reference != nil
}

// Reference returns the reference, if any
func (obj *element) Reference() references.Reference {
	return obj.reference
}

// IsRule returns true if there is a rule, false otherwise
func (obj *element) IsRule() bool {
	return obj.rule != ""
}

// Rule returns the rule, if any
func (obj *element) Rule() string {
	return obj.rule
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *element) IsConstant() bool {
	return obj.constant != ""
}

// Constant returns the constant, if any
func (obj *element) Constant() string {
	return obj.constant
}

// IsBlock returns true if there is a block, false otherwise
func (obj *element) IsBlock() bool {
	return obj.block != ""
}

// Block returns the block, if any
func (obj *element) Block() string {
	return obj.block
}
