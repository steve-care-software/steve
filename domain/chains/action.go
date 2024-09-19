package chains

import "github.com/steve-care-software/steve/domain/hash"

type action struct {
	hash        hash.Hash
	interpreter Interpreter
	transpile   Transpile
}

func createActionWithInterpreter(
	hash hash.Hash,
	interpreter Interpreter,
) Action {
	return createActionInternally(hash, interpreter, nil)
}

func createActionWithTranspile(
	hash hash.Hash,
	transpile Transpile,
) Action {
	return createActionInternally(hash, nil, transpile)
}

func createActionInternally(
	hash hash.Hash,
	interpreter Interpreter,
	transpile Transpile,
) Action {
	out := action{
		hash:        hash,
		interpreter: interpreter,
		transpile:   transpile,
	}

	return &out
}

// Hash returns the hash
func (obj *action) Hash() hash.Hash {
	return obj.hash
}

// IsInterpret returns true if there is an interpreter, false otherwise
func (obj *action) IsInterpret() bool {
	return obj.interpreter != nil
}

// Interpret returns the interpreter, if any
func (obj *action) Interpret() Interpreter {
	return obj.interpreter
}

// IsTranspile returns true if there is a transpile, false otherwise
func (obj *action) IsTranspile() bool {
	return obj.transpile != nil
}

// Transpile returns the transpile, if any
func (obj *action) Transpile() Transpile {
	return obj.transpile
}
