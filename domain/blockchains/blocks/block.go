package blocks

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents"
	"github.com/steve-care-software/steve/domain/hash"
)

type block struct {
	hash       hash.Hash
	content    contents.Content
	result     []byte
	difficulty uint8
}

func createBlock(
	hash hash.Hash,
	content contents.Content,
	result []byte,
	difficulty uint8,
) Block {
	out := block{
		hash:       hash,
		content:    content,
		result:     result,
		difficulty: difficulty,
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

// Difficulty returns the difficulty
func (obj *block) Difficulty() uint8 {
	return obj.difficulty
}
