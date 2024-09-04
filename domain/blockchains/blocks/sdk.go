package blocks

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents"
	"github.com/steve-care-software/steve/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represnets the builder
type Builder interface {
	Create() Builder
	WithContent(content contents.Content) Builder
	WithResult(result []byte) Builder
	WithDifficulty(difficulty uint8) Builder
	Now() (Block, error)
}

// Block represents a block
type Block interface {
	Hash() hash.Hash
	Content() contents.Content
	Result() []byte
	Difficulty() uint8
}
