package constants

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/programs/grammars/blocks/suites"
	"github.com/steve-care-software/steve/domain/scripts/specifics/grammars/constants/tokens"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewConstantBuilder creates a new constant builder
func NewConstantBuilder() ConstantBuilder {
	hashAdapter := hash.NewAdapter()
	return createConstantBuilder(
		hashAdapter,
	)
}

// Builder represents the constants builder
type Builder interface {
	Create() Builder
	WithList(list []Constant) Builder
	Now() (Constants, error)
}

// Constants represents constants
type Constants interface {
	Hash() hash.Hash
	List() []Constant
}

// ConstantBuilder represents the constant builder
type ConstantBuilder interface {
	Create() ConstantBuilder
	WithName(name string) ConstantBuilder
	WithTokens(tokens tokens.Tokens) ConstantBuilder
	WithSuites(suites suites.Suites) ConstantBuilder
	Now() (Constant, error)
}

// Constant represents a constant
type Constant interface {
	Hash() hash.Hash
	Name() string
	Tokens() tokens.Tokens
	HasSuites() bool
	Suites() suites.Suites
}
