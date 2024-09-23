package references

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/grammars/blocks/lines/tokens/elements/references/values"
)

type reference struct {
	hash    hash.Hash
	grammar string
	value   values.Value
}

func createReference(
	hash hash.Hash,
	grammar string,
	value values.Value,
) Reference {
	out := reference{
		hash:    hash,
		grammar: grammar,
		value:   value,
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

// Value returns the value
func (obj *reference) Value() values.Value {
	return obj.value
}
