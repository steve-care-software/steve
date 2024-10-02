package applications

import (
	"github.com/steve-care-software/steve/parsers/domain/asts"
	"github.com/steve-care-software/steve/parsers/domain/asts/instructions"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
)

// SyscallFn represents the syscall func
type SyscallFn func(map[string][]byte) error

// NewApplication creates a new application
func NewApplication() Application {
	elementsAdapter := instructions.NewElementsAdapter()
	astAdapter := asts.NewAdapter()
	return createApplication(
		elementsAdapter,
		astAdapter,
	)
}

// Application represents the interpreter application
type Application interface {
	// Suites executes all the test suites of the grammar
	Suites(grammar grammars.Grammar) error
}
