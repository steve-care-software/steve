package elements

import "github.com/steve-care-software/steve/commons/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	hashAdapter := hash.NewAdapter()
	return createElementBuilder(
		hashAdapter,
	)
}

// Builder represents an elments list
type Builder interface {
	Create() Builder
	WithList(list []Element) Builder
	Now() (Elements, error)
}

// Elements represents elements
type Elements interface {
	List() []Element
	Fetch(name string) (Element, error)
}

// ElementBuilder represents the element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithRule(rule string) ElementBuilder
	WithBlock(block string) ElementBuilder
	WithConstant(constant string) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Hash() hash.Hash
	Name() string
	IsRule() bool
	Rule() string
	IsBlock() bool
	Block() string
	IsConstant() bool
	Constant() string
}
