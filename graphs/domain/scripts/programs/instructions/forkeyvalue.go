package instructions

import "github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/assignments/assignables"

type forKeyValue struct {
	key          string
	value        string
	iterable     assignables.Iterable
	instructions ForInstructions
}

func createForKeyValue(
	key string,
	value string,
	iterable assignables.Iterable,
	instructions ForInstructions,
) ForKeyValue {
	return &forKeyValue{
		key:          key,
		value:        value,
		iterable:     iterable,
		instructions: instructions,
	}
}

// Key returns the key
func (obj *forKeyValue) Key() string {
	return obj.key
}

// Value returns the value
func (obj *forKeyValue) Value() string {
	return obj.value
}

// Iterable returns the iterable
func (obj *forKeyValue) Iterable() assignables.Iterable {
	return obj.iterable
}

// Instructions returns the instructions
func (obj *forKeyValue) Instructions() ForInstructions {
	return obj.instructions
}
