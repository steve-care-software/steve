package scripts

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/bridges"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/contexts"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/grammars"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/pipelines"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/roots"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles"
	"github.com/steve-care-software/steve/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// FetchGrammarInput fetches the grammar input
func FetchGrammarInput() []byte {
	return grammarInput()
}

// ToTranspile converts an input to a script instance
type ParserAdapter interface {
	ToTransfer(input []byte) (Script, []byte, error)
}

// Builder represents a script builder
type Builder interface {
	Create() Builder
	WithGrammar(grammar grammars.Grammar) Builder
	WithTranspile(transpile transpiles.Transpile) Builder
	WithContext(context contexts.Context) Builder
	WithBridge(bridge bridges.Bridge) Builder
	WithProgram(program programs.Program) Builder
	WithPipeline(pipeline pipelines.Pipeline) Builder
	WithRoot(root roots.Root) Builder
	Now() (Script, error)
}

// Script represents a script
type Script interface {
	Hash() hash.Hash
	IsGrammar() bool
	Grammar() grammars.Grammar
	IsTranspile() bool
	Transpile() transpiles.Transpile
	IsContext() bool
	Context() contexts.Context
	IsBridge() bool
	Bridge() bridges.Bridge
	IsProgram() bool
	Program() programs.Program
	IsPipeline() bool
	Pipeline() pipelines.Pipeline
	IsRoot() bool
	Root() roots.Root
}
