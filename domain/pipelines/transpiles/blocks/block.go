package blocks

import "github.com/steve-care-software/steve/domain/pipelines/transpiles/blocks/lines"

type block struct {
	name  string
	lines lines.Lines
}

func createBlock(
	name string,
	lines lines.Lines,
) Block {
	out := block{
		name:  name,
		lines: lines,
	}

	return &out
}

// Name returns the name
func (obj *block) Name() string {
	return obj.name
}

// Lines returns the lines
func (obj *block) Lines() lines.Lines {
	return obj.lines
}
