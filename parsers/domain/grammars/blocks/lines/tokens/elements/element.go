package elements

import "github.com/steve-care-software/steve/hash"

type element struct {
	hash     hash.Hash
	rule     string
	block    string
	constant string
}

func createElementWithRule(hash hash.Hash, rule string) Element {
	return createElementInternally(hash, rule, "", "")
}

func createElementWithBlock(hash hash.Hash, block string) Element {
	return createElementInternally(hash, "", block, "")
}

func createElementWithConstant(hash hash.Hash, constant string) Element {
	return createElementInternally(hash, "", "", constant)
}

func createElementInternally(
	hash hash.Hash,
	rule string,
	block string,
	constant string,
) Element {
	out := element{
		hash:     hash,
		rule:     rule,
		block:    block,
		constant: constant,
	}

	return &out
}

// Hash returns the hash
func (obj *element) Hash() hash.Hash {
	return obj.hash
}

// Name returns the rule or block name
func (obj *element) Name() string {
	if obj.IsBlock() {
		return obj.block
	}

	if obj.IsConstant() {
		return obj.constant
	}

	return obj.rule
}

// IsRule returns true if there is a rule, false otherwise
func (obj *element) IsRule() bool {
	return obj.rule != ""
}

// Rule returns the rule, if any
func (obj *element) Rule() string {
	return obj.rule
}

// IsBlock returns true if there is a block, false otherwise
func (obj *element) IsBlock() bool {
	return obj.block != ""
}

// Block returns the block, if any
func (obj *element) Block() string {
	return obj.block
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *element) IsConstant() bool {
	return obj.constant != ""
}

// Constant returns the constant, if any
func (obj *element) Constant() string {
	return obj.constant
}
