package instructions

import "github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"

type element struct {
	constant    Constant
	instruction Instruction
}

func createElementWithConstant(constant Constant) Element {
	return createElementInternally(constant, nil)
}

func createElementWithInstruction(instruction Instruction) Element {
	return createElementInternally(nil, instruction)
}

func createElementInternally(
	constant Constant,
	instruction Instruction,
) Element {
	out := element{
		constant:    constant,
		instruction: instruction,
	}

	return &out
}

// Name returns the name
func (obj *element) Name() string {
	if obj.IsConstant() {
		return obj.constant.Name()
	}

	return obj.instruction.Block()
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *element) IsConstant() bool {
	return obj.constant != nil
}

// Constant returns the constant, if any
func (obj *element) Constant() Constant {
	return obj.constant
}

// IsInstruction returns true if there is an instruction, false otherwise
func (obj *element) IsInstruction() bool {
	return obj.instruction != nil
}

// Instruction returns the instruction, if any
func (obj *element) Instruction() Instruction {
	return obj.instruction
}

// IsChainValid validates the element against the chain
func (obj *element) IsChainValid(chain chains.Chain) bool {
	if obj.IsInstruction() {
		return obj.instruction.Tokens().IsChainValid(chain)
	}

	return obj.constant.IsChainValid(chain)
}
