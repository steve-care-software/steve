package tokens

import (
	"errors"

	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens/pointers"
	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens/updates"
)

type tokenBuilder struct {
	update updates.Update
	del    pointers.Pointer
	insert pointers.Pointer
}

func createTokenBuilder() TokenBuilder {
	out := tokenBuilder{
		update: nil,
		del:    nil,
		insert: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder()
}

// WithUpdate adds an update to the builder
func (app *tokenBuilder) WithUpdate(update updates.Update) TokenBuilder {
	app.update = update
	return app
}

// WithDelete adds a delete to the builder
func (app *tokenBuilder) WithDelete(delete pointers.Pointer) TokenBuilder {
	app.del = delete
	return app
}

// WithInsert adds an insert to the builder
func (app *tokenBuilder) WithInsert(insert pointers.Pointer) TokenBuilder {
	app.insert = insert
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.update != nil {
		return createTokenWithUpdate(app.update), nil
	}

	if app.del != nil {
		return createTokenWithDelete(app.del), nil
	}

	if app.insert != nil {
		return createTokenWithInsert(app.insert), nil
	}

	return nil, errors.New("the Token is invalid")
}
