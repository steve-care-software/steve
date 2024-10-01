package asts

import (
	"github.com/steve-care-software/steve/parsers/domain/asts/instructions"
)

type ast struct {
	root instructions.Element
}

func createAST(
	root instructions.Element,
) AST {
	return createASTInternally(root)
}

func createASTInternally(
	root instructions.Element,
) AST {
	out := ast{
		root: root,
	}

	return &out
}

// Root returns the root
func (obj *ast) Root() instructions.Element {
	return obj.root
}
