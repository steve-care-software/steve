package instructions

import (
	"errors"
)

type elementBuilder struct {
	constant    Constant
	instruction Instruction
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		constant:    nil,
		instruction: nil,
	}

	return &out
}

// Create initializes the elementBuilder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithConstant adds a constant to the elementBuilder
func (app *elementBuilder) WithConstant(constant Constant) ElementBuilder {
	app.constant = constant
	return app
}

// WithInstruction adds an instruction to the elementBuilder
func (app *elementBuilder) WithInstruction(instruction Instruction) ElementBuilder {
	app.instruction = instruction
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.constant != nil {
		return createElementWithConstant(app.constant), nil
	}

	if app.instruction != nil {
		return createElementWithInstruction(app.instruction), nil
	}

	return nil, errors.New("the Element is invalid")
}
