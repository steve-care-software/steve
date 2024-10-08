package chains

import (
	"github.com/steve-care-software/steve/hash"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/suites"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewActionBuilder creates a new action builder
func NewActionBuilder() ActionBuilder {
	return createActionBuilder()
}

// NewInterpreterBuilder creates a new interpreter builder
func NewInterpreterBuilder() InterpreterBuilder {
	return createInterpreterBuilder()
}

// NewTranspileBuilder creates a new transpile builder
func NewTranspileBuilder() TranspileBuilder {
	return createTranspileBuilder()
}

// Adapter represents a chain adapter
type Adapter interface {
	ToNFT(ins Chain) (Chain, error)
	ToInstance(nft hash.Hash) (Chain, error)
}

// Builder represents a chain builder
type Builder interface {
	Create() Builder
	WithGrammar(grammar hash.Hash) Builder
	WithAction(action Action) Builder
	WithSuites(suites suites.Suites) Builder
	Now() (Chain, error)
}

// Chain represents a chain of action
type Chain interface {
	Grammar() hash.Hash
	Action() Action
	HasSuites() bool
	Suites() suites.Suites
}

// ActionBuilder represents the action builder
type ActionBuilder interface {
	Create() ActionBuilder
	WithInterpret(interpret Interpreter) ActionBuilder
	WithTranspile(transpile Transpile) ActionBuilder
	Now() (Action, error)
}

// Action represents a program action
type Action interface {
	IsInterpret() bool
	Interpret() Interpreter
	IsTranspile() bool
	Transpile() Transpile
}

// InterpreterBuilder represents the interpreter builder
type InterpreterBuilder interface {
	Create() InterpreterBuilder
	WithVariable(variable string) InterpreterBuilder
	WithNext(next Chain) InterpreterBuilder
	Now() (Interpreter, error)
}

// Interpreter represents the interpreter
type Interpreter interface {
	Variable() string
	HasNext() bool
	Next() Chain
}

// TranspileBuilder represents a transpile builder
type TranspileBuilder interface {
	Create() TranspileBuilder
	WithBridge(bridge hash.Hash) TranspileBuilder
	WithTarget(target hash.Hash) TranspileBuilder
	WithNext(next Chain) TranspileBuilder
	Now() (Transpile, error)
}

// Transpile represents a transpile
type Transpile interface {
	Bridge() hash.Hash // transpile code
	Target() hash.Hash // grammar code
	HasNext() bool
	Next() Chain
}
