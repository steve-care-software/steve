package instructions

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewInstructionBuilder creates a new instruction builder
func NewInstructionBuilder() InstructionBuilder {
	return createInstructionBuilder()
}

// NewTokensBuilder creates a new tokens builder
func NewTokensBuilder() TokensBuilder {
	return createTokensBuilder()
}

// NewTokenBuilder creates a new token builder
func NewTokenBuilder() TokenBuilder {
	return createTokenBuilder()
}

// NewElementsAdapter creates a new elements adapter
func NewElementsAdapter() ElementsAdapter {
	return createElementsAdapter()
}

// NewElementsBuilder creates a new elements builder
func NewElementsBuilder() ElementsBuilder {
	return createElementsBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// NewConstantBuilder creates a new constant builder
func NewConstantBuilder() ConstantBuilder {
	return createConstantBuilder()
}

// Builder represents the instructions builder
type Builder interface {
	Create() Builder
	WithList(list []Instruction) Builder
	Now() (Instructions, error)
}

// Instructions represents instructions
type Instructions interface {
	List() []Instruction
	Fetch(name string, idx uint) (Instruction, error)
}

// InstructionBuilder represents the instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithBlock(block string) InstructionBuilder
	WithLine(line uint) InstructionBuilder
	WithTokens(tokens Tokens) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	Block() string
	Line() uint
	Tokens() Tokens
}

// TokensBuilder represents the tokens builder
type TokensBuilder interface {
	Create() TokensBuilder
	WithList(list []Token) TokensBuilder
	Now() (Tokens, error)
}

// Tokens represents tokens
type Tokens interface {
	List() []Token
	Fetch(name string, index uint) (Token, error)
}

// TokenBuilder represents the token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithName(name string) TokenBuilder
	WithElements(elements Elements) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Name() string
	Elements() Elements
}

// ElementsAdapter represents the elements adapter
type ElementsAdapter interface {
	// ToBytes takes an elements and returns its bytes
	ToBytes(elements Elements) ([]byte, error)
}

// ElementsBuilder represents the elements builder
type ElementsBuilder interface {
	Create() ElementsBuilder
	WithList(list []Element) ElementsBuilder
	Now() (Elements, error)
}

// Elements represents elements
type Elements interface {
	List() []Element
}

// ElementBuilder represents the element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithConstant(constant Constant) ElementBuilder
	WithInstruction(instruction Instruction) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Name() string
	IsConstant() bool
	Constant() Constant
	IsInstruction() bool
	Instruction() Instruction
}

// ConstantBuilder represents the constant builder
type ConstantBuilder interface {
	Create() ConstantBuilder
	WithName(name string) ConstantBuilder
	WithValue(value []byte) ConstantBuilder
	Now() (Constant, error)
}

// Constant represents a constant
type Constant interface {
	Name() string
	Value() []byte
}
