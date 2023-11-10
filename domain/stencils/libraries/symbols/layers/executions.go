package layers

import "github.com/steve-care-software/steve/domain/blockchains/hash"

type executions struct {
	hash hash.Hash
	list []Execution
}

func createExecutions(
	hash hash.Hash,
	list []Execution,
) Executions {
	out := executions{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *executions) Hash() hash.Hash {
	return obj.hash
}

// List returns the executions
func (obj *executions) List() []Execution {
	return obj.list
}
