package chains

import (
	"github.com/steve-care-software/steve/domain/chains/nfts"
	"github.com/steve-care-software/steve/domain/hash"
)

// Adapter represents a chain adapterg
type Adapter interface {
	ToNFT(ins Chain) (Chain, error)
	ToInstance(nft nfts.NFT) (Chain, error)
}

// Builder represents a chain builder
type Builder interface {
	Create() Builder
	WithGrammar(grammar nfts.NFT) Builder
	WithProgram(input nfts.NFT) Builder
	WithAction(action Action) Builder
	Now() (Chain, error)
}

// Chain represents a chain of action
type Chain interface {
	Hash() hash.Hash
	Grammar() nfts.NFT // contain my grammar code
	Program() nfts.NFT // contains my program code
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
	Hash() hash.Hash
	IsInterpret() bool
	Interpret() Next
	IsTranspile() bool
	Transpile() Transpile
}

// TranspileBuilder represents a transpile builder
type TranspileBuilder interface {
	Create() TranspileBuilder
	WithBridge(bridge nfts.NFT) TranspileBuilder
	WithTarget(target nfts.NFT) TranspileBuilder
	WithNext(next Next) TranspileBuilder
	Now() (Transpile, error)
}

// Transpile represents a transpile
type Transpile interface {
	Hash() hash.Hash
	Bridge() nfts.NFT // bridge code
	Target() nfts.NFT // grammar code
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
	Hash() hash.Hash
	IsOutput() bool
	IsChain() bool
	Chain() Chain
}
