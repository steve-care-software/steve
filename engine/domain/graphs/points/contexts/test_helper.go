package contexts

import (
	"github.com/steve-care-software/steve/engine/domain/hash"
)

// NewContextsForTests creates new contexts for tests
func NewContextsForTests(list []Context) Contexts {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContextForTests creates a new context for tests
func NewContextForTests(name string) Context {
	ins, err := NewContextBuilder().Create().WithName(name).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContextWithParentForTests creates a new context with parent for tests
func NewContextWithParentForTests(name string, parent hash.Hash) Context {
	ins, err := NewContextBuilder().Create().WithName(name).WithParent(parent).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
