package blocks

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/blockchains/blocks/contents"
)

type block struct {
	hash    hash.Hash
	content contents.Content
	result  []byte
}

func createBlock(
	hash hash.Hash,
	content contents.Content,
	result []byte,
) Block {
	out := block{
		hash:    hash,
		content: content,
		result:  result,
	}

	return &out
}

// Hash returns the hash
func (obj *block) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *block) Content() contents.Content {
	return obj.content
}

// Result returns the result
func (obj *block) Result() []byte {
	return obj.result
}
