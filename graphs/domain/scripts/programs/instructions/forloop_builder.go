package instructions

import "errors"

type forLoopBuilder struct {
	index    ForIndex
	keyValue ForKeyValue
}

func createForLoopBuilder() ForLoopBuilder {
	return &forLoopBuilder{
		index:    nil,
		keyValue: nil,
	}
}

func (app *forLoopBuilder) Create() ForLoopBuilder {
	return createForLoopBuilder()
}

func (app *forLoopBuilder) WithIndex(index ForIndex) ForLoopBuilder {
	app.index = index
	return app
}

func (app *forLoopBuilder) WithKeyValue(keyValue ForKeyValue) ForLoopBuilder {
	app.keyValue = keyValue
	return app
}

func (app *forLoopBuilder) Now() (ForLoop, error) {
	if app.index != nil {
		return createForLoopWithIndex(app.index), nil
	}

	if app.keyValue != nil {
		return createForLoopWithKeyValue(app.keyValue), nil
	}

	return nil, errors.New("the ForLoop is invalid")
}
