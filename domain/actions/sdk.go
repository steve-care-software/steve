package actions

import "github.com/steve-care-software/steve/domain/actions/nfts"

// Builder represents a chain builder
type Builder interface {
	Create() Builder
	WithInput(input []byte) Builder
	WithGrammar(grammar []byte) Builder
	WithAction(action Action) Builder
	Now() (Chain, error)
}

// Chain represents a chain of action
type Chain interface {
	Input() nfts.NFT   // contains my program code
	Grammar() nfts.NFT // contain my grammar code
	Action() Action
}

// ActionBuilder represents the action builder
type ActionBuilder interface {
	Create() ActionBuilder
	WithInterpret(interpret Next) ActionBuilder
	WithTranspile(transpile Transpile) ActionBuilder
	Now() (Action, error)
}

// Action represents a program action
type Action interface {
	IsInterpret() bool
	Interpret() Next
	IsTranspile() bool
	Transpile() Transpile
}

// TranspileBuilder represents a transpile builder
type TranspileBuilder interface {
	Create() TranspileBuilder
	WithSource(source []byte) TranspileBuilder
	WithBridge(bridge []byte) TranspileBuilder
	WithNext(next Next) TranspileBuilder
	Now() (Transpile, error)
}

// Transpile represents a transpile
type Transpile interface {
	Source() nfts.NFT // grammar code
	Bridge() nfts.NFT // bridge code
	HasNext() bool
	Next() Next
}

// NextBuilder represents the next builder
type NextBuilder interface {
	Create() NextBuilder
	WithChain(chain Chain) NextBuilder
	IsOutput() NextBuilder
	Now() (Next, error)
}

// Next represents the next step
type Next interface {
	IsOutput() bool
	IsChain() bool
	Chain() Chain
}
