package transpiles

import (
	"github.com/steve-care-software/steve/domain/pipelines/transpiles/blocks"
	"github.com/steve-care-software/steve/domain/pipelines/transpiles/blocks/lines"
	"github.com/steve-care-software/steve/domain/pipelines/transpiles/blocks/lines/tokens"
	"github.com/steve-care-software/steve/domain/pipelines/transpiles/blocks/lines/tokens/pointers"
	"github.com/steve-care-software/steve/domain/pipelines/transpiles/blocks/lines/tokens/pointers/elements"
	"github.com/steve-care-software/steve/domain/pipelines/transpiles/blocks/lines/tokens/updates"
	"github.com/steve-care-software/steve/domain/pipelines/transpiles/blocks/lines/tokens/updates/targets"
	"github.com/steve-care-software/steve/domain/programs/grammars"
)

type parserAdapter struct {
	grammarParserAdapter grammars.ParserAdapter
	transpileBuilder     Builder
	blocksBuilder        blocks.Builder
	blockBuilder         blocks.BlockBuilder
	linesBuilder         lines.Builder
	lineBuilder          lines.LineBuilder
	tokensBuilder        tokens.Builder
	tokenBuilder         tokens.TokenBuilder
	updateBuilder        updates.Builder
	targetBuilder        targets.Builder
	pointerBuilder       pointers.Builder
	elementBuilder       elements.Builder
}

func createParserAdapter(
	grammarParserAdapter grammars.ParserAdapter,
	transpileBuilder Builder,
	blocksBuilder blocks.Builder,
	blockBuilder blocks.BlockBuilder,
	linesBuilder lines.Builder,
	lineBuilder lines.LineBuilder,
	tokensBuilder tokens.Builder,
	tokenBuilder tokens.TokenBuilder,
	updateBuilder updates.Builder,
	targetBuilder targets.Builder,
	pointerBuilder pointers.Builder,
	elementBuilder elements.Builder,
) ParserAdapter {
	out := parserAdapter{
		grammarParserAdapter: grammarParserAdapter,
		transpileBuilder:     transpileBuilder,
		blocksBuilder:        blocksBuilder,
		blockBuilder:         blockBuilder,
		linesBuilder:         linesBuilder,
		lineBuilder:          lineBuilder,
		tokensBuilder:        tokensBuilder,
		tokenBuilder:         tokenBuilder,
		updateBuilder:        updateBuilder,
		targetBuilder:        targetBuilder,
		pointerBuilder:       pointerBuilder,
		elementBuilder:       elementBuilder,
	}

	return &out
}

// ToTranspile converts an input to a transpile instance
func (app *parserAdapter) ToTranspile(input []byte) (Transpile, []byte, error) {
	return nil, nil, nil
}
