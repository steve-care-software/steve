package chains

import (
	"github.com/steve-care-software/steve/domain/chains/nfts"
	"github.com/steve-care-software/steve/domain/hash"
)

// NewActionBuilder creates a new action builder
func NewActionBuilder() ActionBuilder {
	hashAdapter := hash.NewAdapter()
	return createActionBuilder(
		hashAdapter,
	)
}

// NewInterpreterBuilder creates a new interpreter builder
func NewInterpreterBuilder() InterpreterBuilder {
	hashAdapter := hash.NewAdapter()
	return createInterpreterBuilder(
		hashAdapter,
	)
}

// NewTranspileBuilder creates a new transpile builder
func NewTranspileBuilder() TranspileBuilder {
	hashAdapter := hash.NewAdapter()
	return createTranspileBuilder(
		hashAdapter,
	)
}

// Adapter represents a chain adapterg
type Adapter interface {
	ToNFT(ins Chain) (Chain, error)
	ToInstance(nft nfts.NFT) (Chain, error)
}

// Builder represents a chain builder
type Builder interface {
	Create() Builder
	WithGrammar(grammar nfts.NFT) Builder
	WithAction(action Action) Builder
	Now() (Chain, error)
}

// Chain represents a chain of action
type Chain interface {
	Hash() hash.Hash
	Grammar() nfts.NFT // contain my grammar code
	Action() Action
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
	Hash() hash.Hash
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
	Hash() hash.Hash
	Variable() string
	HasNext() bool
	Next() Chain
}

// TranspileBuilder represents a transpile builder
type TranspileBuilder interface {
	Create() TranspileBuilder
	WithBridge(bridge nfts.NFT) TranspileBuilder
	WithTarget(target nfts.NFT) TranspileBuilder
	WithNext(next Chain) TranspileBuilder
	Now() (Transpile, error)
}

// Transpile represents a transpile
type Transpile interface {
	Hash() hash.Hash
	Bridge() nfts.NFT // bridge code
	Target() nfts.NFT // grammar code
	HasNext() bool
	Next() Chain
}
