package blocks

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents"
	"github.com/steve-care-software/steve/domain/hash"
)

// Block represents a block
type Block interface {
	Hash() hash.Hash
	Content() contents.Content
	Result() []byte
	Difficulty() uint
}
