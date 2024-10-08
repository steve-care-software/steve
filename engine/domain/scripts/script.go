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

type script struct {
	hash      hash.Hash
	grammar   grammars.Grammar
	transpile transpiles.Transpile
	context   contexts.Context
	bridge    bridges.Bridge
	program   programs.Program
	pipeline  pipelines.Pipeline
	root      roots.Root
}

func createScriptWithGrammar(
	hash hash.Hash,
	grammar grammars.Grammar,
) Script {
	return createScriptInternally(
		hash,
		grammar,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createScriptWithTranspile(
	hash hash.Hash,
	transpile transpiles.Transpile,
) Script {
	return createScriptInternally(
		hash,
		nil,
		transpile,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createScriptWithContext(
	hash hash.Hash,
	context contexts.Context,
) Script {
	return createScriptInternally(
		hash,
		nil,
		nil,
		context,
		nil,
		nil,
		nil,
		nil,
	)
}

func createScriptWithBridge(
	hash hash.Hash,
	bridge bridges.Bridge,
) Script {
	return createScriptInternally(
		hash,
		nil,
		nil,
		nil,
		bridge,
		nil,
		nil,
		nil,
	)
}

func createScriptWithProgram(
	hash hash.Hash,
	program programs.Program,
) Script {
	return createScriptInternally(
		hash,
		nil,
		nil,
		nil,
		nil,
		program,
		nil,
		nil,
	)
}

func createScriptWithPipeline(
	hash hash.Hash,
	pipeline pipelines.Pipeline,
) Script {
	return createScriptInternally(
		hash,
		nil,
		nil,
		nil,
		nil,
		nil,
		pipeline,
		nil,
	)
}

func createScriptWithRoot(
	hash hash.Hash,
	root roots.Root,
) Script {
	return createScriptInternally(
		hash,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		root,
	)
}

func createScriptInternally(
	hash hash.Hash,
	grammar grammars.Grammar,
	transpile transpiles.Transpile,
	context contexts.Context,
	bridge bridges.Bridge,
	program programs.Program,
	pipeline pipelines.Pipeline,
	root roots.Root,
) Script {
	out := script{
		hash:      hash,
		grammar:   grammar,
		transpile: transpile,
		context:   context,
		bridge:    bridge,
		program:   program,
		pipeline:  pipeline,
		root:      root,
	}

	return &out
}

// Hash returns the hash
func (obj *script) Hash() hash.Hash {
	return obj.hash
}

// IsGrammar returns true if there is a grammar, false otherwise
func (obj *script) IsGrammar() bool {
	return obj.grammar != nil
}

// Grammar returns the grammar, if any
func (obj *script) Grammar() grammars.Grammar {
	return obj.grammar
}

// IsTranspile returns true if there is a transpile, false otherwise
func (obj *script) IsTranspile() bool {
	return obj.transpile != nil
}

// Transpile returns the transpile, if any
func (obj *script) Transpile() transpiles.Transpile {
	return obj.transpile
}

// IsContext returns true if there is a context, false otherwise
func (obj *script) IsContext() bool {
	return obj.context != nil
}

// Context returns the context, if any
func (obj *script) Context() contexts.Context {
	return obj.context
}

// IsBridge returns true if there is a bridge, false otherwise
func (obj *script) IsBridge() bool {
	return obj.bridge != nil
}

// Bridge returns the bridge, if any
func (obj *script) Bridge() bridges.Bridge {
	return obj.bridge
}

// IsProgram returns true if there is a program, false otherwise
func (obj *script) IsProgram() bool {
	return obj.program != nil
}

// Program returns the program, if any
func (obj *script) Program() programs.Program {
	return obj.program
}

// IsPipeline returns true if there is a pipeline, false otherwise
func (obj *script) IsPipeline() bool {
	return obj.pipeline != nil
}

// Pipeline returns the pipeline, if any
func (obj *script) Pipeline() pipelines.Pipeline {
	return obj.pipeline
}

// IsRoot returns true if there is a root, false otherwise
func (obj *script) IsRoot() bool {
	return obj.root != nil
}

// Root returns the root, if any
func (obj *script) Root() roots.Root {
	return obj.root
}
