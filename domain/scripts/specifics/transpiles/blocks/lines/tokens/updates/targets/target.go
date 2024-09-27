package targets

import "github.com/steve-care-software/steve/domain/hash"

type target struct {
	hash     hash.Hash
	constant string
	rule     string
}

func createTargetWithConstant(
	hash hash.Hash,
	constant string,
) Target {
	return createTargetInternally(hash, constant, "")
}

func createTargetWithRule(
	hash hash.Hash,
	rule string,
) Target {
	return createTargetInternally(hash, "", rule)
}

func createTargetInternally(
	hash hash.Hash,
	constant string,
	rule string,
) Target {
	out := target{
		hash:     hash,
		constant: constant,
		rule:     rule,
	}

	return &out
}

// Hash returns the hash
func (obj *target) Hash() hash.Hash {
	return obj.hash
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *target) IsConstant() bool {
	return obj.constant != ""
}

// Constant returns the constant, if any
func (obj *target) Constant() string {
	return obj.constant
}

// IsRule returns true if there is a rule, false otherwise
func (obj *target) IsRule() bool {
	return obj.rule != ""
}

// Rule returns the rule, if any
func (obj *target) Rule() string {
	return obj.rule
}
