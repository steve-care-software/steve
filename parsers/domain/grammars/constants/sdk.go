package constants

import "github.com/steve-care-software/steve/parsers/domain/grammars/constants/tokens"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewConstantBuilder creates a new constant builder
func NewConstantBuilder() ConstantBuilder {
	return createConstantBuilder()
}

// Builder represents the constants builder
type Builder interface {
	Create() Builder
	WithList(list []Constant) Builder
	Now() (Constants, error)
}

// Constants represents constants
type Constants interface {
	List() []Constant
	Fetch(name string) (Constant, error)
}

// ConstantBuilder represents the constant builder
type ConstantBuilder interface {
	Create() ConstantBuilder
	WithName(name string) ConstantBuilder
	WithTokens(tokens tokens.Tokens) ConstantBuilder
	Now() (Constant, error)
}

// Constant represents a constant
type Constant interface {
	Name() string
	Tokens() tokens.Tokens
}
