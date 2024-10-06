package walkers

import "errors"

type nodeBuilder struct {
	token     Token
	tokenList TokenList
	element   Walker
}

func createNodeBuilder() NodeBuilder {
	out := nodeBuilder{
		token:     nil,
		tokenList: nil,
		element:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *nodeBuilder) Create() NodeBuilder {
	return createNodeBuilder()
}

// WithToken adds a token to the builder
func (app *nodeBuilder) WithToken(token Token) NodeBuilder {
	app.token = token
	return app
}

// WithTokenList adds a tokenList to the builder
func (app *nodeBuilder) WithTokenList(tokenList TokenList) NodeBuilder {
	app.tokenList = tokenList
	return app
}

// WithElement adds an element to the builder
func (app *nodeBuilder) WithElement(element Walker) NodeBuilder {
	app.element = element
	return app
}

// Now builds a new Node instance
func (app *nodeBuilder) Now() (Node, error) {
	if app.token != nil {
		return createNodeWithToken(app.token), nil
	}

	if app.tokenList != nil {
		return createNodeWithTokenList(app.tokenList), nil
	}

	if app.element != nil {
		return createNodeWithElement(app.element), nil
	}

	return nil, errors.New("the Node is invalid")
}
