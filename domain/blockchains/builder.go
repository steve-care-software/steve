package blockchains

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/domain/blockchains/roots"
	"github.com/steve-care-software/steve/domain/blockchains/rules"
)

type builder struct {
	pIdentifier *uuid.UUID
	name        string
	description string
	rules       rules.Rules
	root        roots.Root
	head        blocks.Block
	pCreatedOn  *time.Time
}

func createBuilder() Builder {
	out := builder{
		pIdentifier: nil,
		name:        "",
		description: "",
		rules:       nil,
		root:        nil,
		head:        nil,
		pCreatedOn:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithIdentifier adds an identifier to the builder
func (app *builder) WithIdentifier(identifier uuid.UUID) Builder {
	app.pIdentifier = &identifier
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithDescription adds a description to the builder
func (app *builder) WithDescription(description string) Builder {
	app.description = description
	return app
}

// WithRules add rules to the builder
func (app *builder) WithRules(rules rules.Rules) Builder {
	app.rules = rules
	return app
}

// WithRoot add root to the builder
func (app *builder) WithRoot(root roots.Root) Builder {
	app.root = root
	return app
}

// WithHead add head to the builder
func (app *builder) WithHead(head blocks.Block) Builder {
	app.head = head
	return app
}

// CreatedOn add creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Blockchain instance
func (app *builder) Now() (Blockchain, error) {
	if app.pIdentifier == nil {
		return nil, errors.New("the identifier is mandatory in order to build a Blockchain instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Blockchain instance")
	}

	if app.rules == nil {
		return nil, errors.New("the rules is mandatory in order to build a Blockchain instance")
	}

	if app.root == nil {
		return nil, errors.New("the root is mandatory in order to build a Blockchain instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Blockchain instance")
	}

	if app.head != nil {
		return createBlockchainWithHead(
			*app.pIdentifier,
			app.name,
			app.description,
			app.rules,
			app.root,
			*app.pCreatedOn,
			app.head,
		), nil
	}

	return createBlockchain(
		*app.pIdentifier,
		app.name,
		app.description,
		app.rules,
		app.root,
		*app.pCreatedOn,
	), nil
}
