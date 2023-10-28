package kinds

import "github.com/steve-care-software/steve/domain/hash"

type kind struct {
	hash       hash.Hash
	isContinue bool
	isPrompt   bool
	execute    []string
}

func createKindWithContinue(
	hash hash.Hash,
) Kind {
	return createKindInternally(hash, true, false, nil)
}

func createKindWithPrompt(
	hash hash.Hash,
) Kind {
	return createKindInternally(hash, false, true, nil)
}

func createKindWithExecute(
	hash hash.Hash,
	execute []string,
) Kind {
	return createKindInternally(hash, false, false, execute)
}

func createKindInternally(
	hash hash.Hash,
	isContinue bool,
	isPrompt bool,
	execute []string,
) Kind {
	out := kind{
		hash:       hash,
		isContinue: isContinue,
		isPrompt:   isPrompt,
		execute:    execute,
	}

	return &out
}

// Hash returns the hash
func (obj *kind) Hash() hash.Hash {
	return obj.hash
}

// IsContinue returns true if continue, false otherwise
func (obj *kind) IsContinue() bool {
	return obj.isContinue
}

// IsPrompt returns true if prompt, false otherwise
func (obj *kind) IsPrompt() bool {
	return obj.isPrompt
}

// IsExecute returns true if execute, false otherwise
func (obj *kind) IsExecute() bool {
	return obj.execute != nil
}

// Execute returns the execute commands, if any
func (obj *kind) Execute() []string {
	return obj.execute
}
