package tokens

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles/blocks/lines/tokens/pointers"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles/blocks/lines/tokens/updates"
	"github.com/steve-care-software/steve/hash"
)

type tokenBuilder struct {
	hashAdapter hash.Adapter
	update      updates.Update
	insert      pointers.Pointer
}

func createTokenBuilder(
	hashAdapter hash.Adapter,
) TokenBuilder {
	out := tokenBuilder{
		hashAdapter: hashAdapter,
		update:      nil,
		insert:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder(
		app.hashAdapter,
	)
}

// WithUpdate adds an update to the builder
func (app *tokenBuilder) WithUpdate(update updates.Update) TokenBuilder {
	app.update = update
	return app
}

// WithInsert adds an insert to the builder
func (app *tokenBuilder) WithInsert(insert pointers.Pointer) TokenBuilder {
	app.insert = insert
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	data := [][]byte{}
	if app.update != nil {
		data = append(data, app.update.Hash().Bytes())
	}

	if app.insert != nil {
		data = append(data, app.insert.Hash().Bytes())
	}

	if len(data) != 1 {
		return nil, errors.New("the Token is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.update != nil {
		return createTokenWithUpdate(*pHash, app.update), nil
	}

	return createTokenWithInsert(*pHash, app.insert), nil
}
