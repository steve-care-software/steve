package chains

import "errors"

type elementBuilder struct {
	pIndex *uint
	chain  Chain
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		pIndex: nil,
		chain:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithIndex adds an index to the builder
func (app *elementBuilder) WithIndex(index uint) ElementBuilder {
	app.pIndex = &index
	return app
}

// WithChain adds a chain to the builder
func (app *elementBuilder) WithChain(chain Chain) ElementBuilder {
	app.chain = chain
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build an Element instance")
	}

	if app.chain != nil {
		return createElementWithChain(
			*app.pIndex,
			app.chain,
		), nil
	}

	return createElement(
		*app.pIndex,
	), nil
}
