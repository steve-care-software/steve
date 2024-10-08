package blocks

import (
	"github.com/steve-care-software/steve/blockchains/domain/blocks/contents"
	"github.com/steve-care-software/steve/hash"
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
