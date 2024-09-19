package chains

import "github.com/steve-care-software/steve/domain/hash"

type interpreter struct {
	hash     hash.Hash
	variable string
	next     Chain
}

func createInterpreter(
	hash hash.Hash,
	variable string,
) Interpreter {
	return createInterpreterInternally(hash, variable, nil)
}

func createInterpreterWithNext(
	hash hash.Hash,
	variable string,
	next Chain,
) Interpreter {
	return createInterpreterInternally(hash, variable, next)
}

func createInterpreterInternally(
	hash hash.Hash,
	variable string,
	next Chain,
) Interpreter {
	out := interpreter{
		hash:     hash,
		variable: variable,
		next:     next,
	}

	return &out
}

// Hash returns the hash
func (obj *interpreter) Hash() hash.Hash {
	return obj.hash
}

// Variable returns the variable
func (obj *interpreter) Variable() string {
	return obj.variable
}

// HasNext returns true if next, false otherwise
func (obj *interpreter) HasNext() bool {
	return obj.next != nil
}

// Next returns the next, if any
func (obj *interpreter) Next() Chain {
	return obj.next
}
