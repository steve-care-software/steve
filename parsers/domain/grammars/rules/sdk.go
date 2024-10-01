package rules

import "github.com/steve-care-software/steve/engine/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewRuleBuilder creates a new rule builder
func NewRuleBuilder() RuleBuilder {
	hashAdapter := hash.NewAdapter()
	return createRuleBuilder(
		hashAdapter,
	)
}

// Builder represents a rule list
type Builder interface {
	Create() Builder
	WithList(list []Rule) Builder
	Now() (Rules, error)
}

// Rules represents rules
type Rules interface {
	Hash() hash.Hash
	List() []Rule
	Fetch(name string) (Rule, error)
}

// RuleBuilder represents a rule builder
type RuleBuilder interface {
	Create() RuleBuilder
	WithName(name string) RuleBuilder
	WithBytes(bytes []byte) RuleBuilder
	Now() (Rule, error)
}

// Rule represents a rule
type Rule interface {
	Hash() hash.Hash
	Name() string
	Bytes() []byte
}
