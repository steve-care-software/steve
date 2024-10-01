package scripts

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/bridges"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/contexts"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/grammars"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/pipelines"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/roots"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/schemas"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles"
)

type builder struct {
	hashAdapter hash.Adapter
	grammar     grammars.Grammar
	transpile   transpiles.Transpile
	schema      schemas.Schema
	context     contexts.Context
	bridge      bridges.Bridge
	program     programs.Program
	pipeline    pipelines.Pipeline
	root        roots.Root
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		grammar:   nil,
		transpile: nil,
		schema:    nil,
		context:   nil,
		bridge:    nil,
		program:   nil,
		pipeline:  nil,
		root:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithGrammar adds a grammar to the builder
func (app *builder) WithGrammar(grammar grammars.Grammar) Builder {
	app.grammar = grammar
	return app
}

// WithTranspile adds a transpile to the builder
func (app *builder) WithTranspile(transpile transpiles.Transpile) Builder {
	app.transpile = transpile
	return app
}

// WithSchema adds a schema to the builder
func (app *builder) WithSchema(schema schemas.Schema) Builder {
	app.schema = schema
	return app
}

// WithContext adds a context to the builder
func (app *builder) WithContext(context contexts.Context) Builder {
	app.context = context
	return app
}

// WithBridge adds a bridge to the builder
func (app *builder) WithBridge(bridge bridges.Bridge) Builder {
	app.bridge = bridge
	return app
}

// WithProgram adds a program to the builder
func (app *builder) WithProgram(program programs.Program) Builder {
	app.program = program
	return app
}

// WithPipeline adds a pipeline to the builder
func (app *builder) WithPipeline(pipeline pipelines.Pipeline) Builder {
	app.pipeline = pipeline
	return app
}

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root roots.Root) Builder {
	app.root = root
	return app
}

// Now builds a new Script instance
func (app *builder) Now() (Script, error) {
	data := [][]byte{}
	if app.grammar != nil {
		data = append(data, app.grammar.Hash().Bytes())
	}

	if app.transpile != nil {
		data = append(data, app.transpile.Hash().Bytes())
	}

	if app.schema != nil {
		data = append(data, app.schema.Hash().Bytes())
	}

	if app.context != nil {
		data = append(data, app.context.Hash().Bytes())
	}

	if app.bridge != nil {
		data = append(data, app.bridge.Hash().Bytes())
	}

	if app.program != nil {
		data = append(data, app.program.Hash().Bytes())
	}

	if app.pipeline != nil {
		data = append(data, app.pipeline.Hash().Bytes())
	}

	if app.root != nil {
		data = append(data, app.root.Hash().Bytes())
	}

	if len(data) != 1 {
		return nil, errors.New("the Script is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.grammar != nil {
		return createScriptWithGrammar(*pHash, app.grammar), nil
	}

	if app.transpile != nil {
		return createScriptWithTranspile(*pHash, app.transpile), nil
	}

	if app.schema != nil {
		return createScriptWithSchema(*pHash, app.schema), nil
	}

	if app.context != nil {
		return createScriptWithContext(*pHash, app.context), nil
	}

	if app.bridge != nil {
		return createScriptWithBridge(*pHash, app.bridge), nil
	}

	if app.program != nil {
		return createScriptWithProgram(*pHash, app.program), nil
	}

	if app.pipeline != nil {
		return createScriptWithPipeline(*pHash, app.pipeline), nil
	}

	return createScriptWithRoot(*pHash, app.root), nil
}
