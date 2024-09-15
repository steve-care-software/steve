package blocks

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents"
	"github.com/steve-care-software/steve/domain/hash"
)

const dataLengthTooSmallErrPattern = "(block) the data length was expected to be at least %d bytes, %d returned"

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	contentAdapter := contents.NewAdapter()
	builder := NewBuilder()
	blockBuilder := NewBlockBuilder()
	return createAdapter(
		contentAdapter,
		builder,
		blockBuilder,
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewBlockBuilder creates a new builder
func NewBlockBuilder() BlockBuilder {
	hashAdapter := hash.NewAdapter()
	return createBlockBuilder(
		hashAdapter,
	)
}

// Adapter represents the block adapter
type Adapter interface {
	InstancesToBytes(ins Blocks) ([]byte, error)
	BytesToInstances(data []byte) (Blocks, []byte, error)
	InstanceToBytes(ins Block) ([]byte, error)
	BytesToInstance(data []byte) (Block, []byte, error)
}

// Builder represents the blocks builder
type Builder interface {
	Create() Builder
	WithList(list []Block) Builder
	Now() (Blocks, error)
}

// Blocks represents the blocks
type Blocks interface {
	Hash() hash.Hash
	List() []Block
}

// BlockBuilder represnets the builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithContent(content contents.Content) BlockBuilder
	WithResult(result []byte) BlockBuilder
	Now() (Block, error)
}

// Block represents a block
type Block interface {
	Hash() hash.Hash
	Content() contents.Content
	Result() []byte
}
