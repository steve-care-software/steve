package asts

import (
	"errors"

	"github.com/steve-care-software/steve/parsers/domain/asts/instructions"
)

type builder struct {
	root instructions.Element
}

func createBuilder() Builder {
	out := builder{
		root: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root instructions.Element) Builder {
	app.root = root
	return app
}

// Now builds a new AST
func (app *builder) Now() (AST, error) {
	if app.root == nil {
		return nil, errors.New("the root is mandatory in order to build a AST instance")
	}

	return createAST(
		app.root,
	), nil
}
