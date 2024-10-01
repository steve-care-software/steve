package values

import "github.com/steve-care-software/steve/engine/domain/hash"

type value struct {
	hash     hash.Hash
	block    string
	constant string
	rule     string
}

func createValueWithBlock(
	hash hash.Hash,
	block string,
) Value {
	return createValueInternally(hash, block, "", "")
}

func createValueWithConstant(
	hash hash.Hash,
	constant string,
) Value {
	return createValueInternally(hash, "", constant, "")
}

func createValueWithRule(
	hash hash.Hash,
	rule string,
) Value {
	return createValueInternally(hash, "", "", rule)
}

func createValueInternally(
	hash hash.Hash,
	block string,
	constant string,
	rule string,
) Value {
	out := value{
		hash:     hash,
		block:    block,
		constant: constant,
		rule:     rule,
	}

	return &out
}

// Hash returns the hash
func (obj *value) Hash() hash.Hash {
	return obj.hash
}

// IsBlock returns true if there is a block, false otherwise
func (obj *value) IsBlock() bool {
	return obj.block != ""
}

// Block returns the block, if any
func (obj *value) Block() string {
	return obj.block
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *value) IsConstant() bool {
	return obj.constant != ""
}

// Constant returns the constant, if any
func (obj *value) Constant() string {
	return obj.constant
}

// IsRule returns true if there is a rule, false otherwise
func (obj *value) IsRule() bool {
	return obj.rule != ""
}

// Rule returns the rule, if any
func (obj *value) Rule() string {
	return obj.rule
}
