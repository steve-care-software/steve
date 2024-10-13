package grammars

import (
	"errors"
	"strings"

	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/grammars/blocks"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/grammars/constants"
	"github.com/steve-care-software/steve/hash"
	"github.com/steve-care-software/steve/parsers/domain/grammars/rules"
)

type builder struct {
	hashAdapter hash.Adapter
	entry       string
	omit        []string
	rules       rules.Rules
	blocks      blocks.Blocks
	constants   constants.Constants
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		entry:       "",
		omit:        nil,
		rules:       nil,
		blocks:      nil,
		constants:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithEntry adds an entry to the builder
func (app *builder) WithEntry(entry string) Builder {
	app.entry = entry
	return app
}

// WithOmit adds an omit to the builder
func (app *builder) WithOmit(omit []string) Builder {
	app.omit = omit
	return app
}

// WithRules add rules to the builder
func (app *builder) WithRules(rules rules.Rules) Builder {
	app.rules = rules
	return app
}

// WithBlocks add blocks to the builder
func (app *builder) WithBlocks(blocks blocks.Blocks) Builder {
	app.blocks = blocks
	return app
}

// WithConstants add constants to the builder
func (app *builder) WithConstants(constants constants.Constants) Builder {
	app.constants = constants
	return app
}

// Now builds a new Grammar instance
func (app *builder) Now() (Grammar, error) {
	if app.entry == "" {
		return nil, errors.New("the entry is mandatory in order to build a Grammar instance")
	}

	if app.rules == nil {
		return nil, errors.New("the rules is mandatory in order to build a Grammar instance")
	}

	if app.blocks == nil {
		return nil, errors.New("the blocks is mandatory in order to build a Grammar instance")
	}

	if app.omit != nil && len(app.omit) <= 0 {
		app.omit = nil
	}

	data := [][]byte{
		[]byte(app.entry),
		app.rules.Hash().Bytes(),
		app.blocks.Hash().Bytes(),
	}

	if app.omit != nil {
		data = append(data, []byte(strings.Join(app.omit, ",")))
	}

	if app.constants != nil {
		data = append(data, app.constants.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.omit != nil && app.constants != nil {
		return createGrammarWithOmitAndConstants(
			*pHash,
			app.entry,
			app.rules,
			app.blocks,
			app.omit,
			app.constants,
		), nil
	}

	if app.omit != nil {
		return createGrammarWithOmit(
			*pHash,
			app.entry,
			app.rules,
			app.blocks,
			app.omit,
		), nil
	}

	if app.constants != nil {
		return createGrammarWithConstants(
			*pHash,
			app.entry,
			app.rules,
			app.blocks,
			app.constants,
		), nil
	}

	return createGrammar(
		*pHash,
		app.entry,
		app.rules,
		app.blocks,
	), nil
}
