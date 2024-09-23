package references

import "github.com/steve-care-software/steve/domain/hash"

type reference struct {
	hash    hash.Hash
	grammar string
	block   string
}

func createReference(
	hash hash.Hash,
	grammar string,
	block string,
) Reference {
	out := reference{
		hash:    hash,
		grammar: grammar,
		block:   block,
	}

	return &out
}

// Hash returns the hash
func (obj *reference) Hash() hash.Hash {
	return obj.hash
}

// Grammar returns the grammar
func (obj *reference) Grammar() string {
	return obj.grammar
}

// Block returns the block
func (obj *reference) Block() string {
	return obj.block
}
