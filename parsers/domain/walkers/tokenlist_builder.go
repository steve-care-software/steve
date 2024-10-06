package walkers

import "errors"

type tokenListBuilder struct {
	fn   MapFn
	list []SelectedTokenList
}

func createTokenListBuilder() TokenListBuilder {
	out := tokenListBuilder{
		fn:   nil,
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenListBuilder) Create() TokenListBuilder {
	return createTokenListBuilder()
}

// WithFn adds a MapFn to the builder
func (app *tokenListBuilder) WithFn(fn MapFn) TokenListBuilder {
	app.fn = fn
	return app
}

// WithList adds a list to the builder
func (app *tokenListBuilder) WithList(list []SelectedTokenList) TokenListBuilder {
	app.list = list
	return app
}

// Now builds a new TokenList instance
func (app *tokenListBuilder) Now() (TokenList, error) {
	if app.fn == nil {
		return nil, errors.New("the MapFn is mandatory in order to build a TokenList instance")
	}

	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 SelectedTokenList in order to build a TokenList instance")
	}

	return createTokenList(
		app.fn,
		app.list,
	), nil
}
