package instructions

import "github.com/steve-care-software/steve/domain/hash"

type loopHeader struct {
	hash       hash.Hash
	counter    LoopCounter
	keyValue   LoopKeyValue
	isInfinite bool
}

func createLoopHeaderWithCounter(
	hash hash.Hash,
	counter LoopCounter,
) LoopHeader {
	return createLoopHeaderInternally(hash, counter, nil, false)
}

func createLoopHeaderWithKeyValue(
	hash hash.Hash,
	keyValue LoopKeyValue,
) LoopHeader {
	return createLoopHeaderInternally(hash, nil, keyValue, false)
}

func createLoopHeaderWithKeyInfinite(
	hash hash.Hash,
) LoopHeader {
	return createLoopHeaderInternally(hash, nil, nil, true)
}

func createLoopHeaderInternally(
	hash hash.Hash,
	counter LoopCounter,
	keyValue LoopKeyValue,
	isInfinite bool,
) LoopHeader {
	out := loopHeader{
		hash:       hash,
		counter:    counter,
		keyValue:   keyValue,
		isInfinite: isInfinite,
	}

	return &out
}

// Hash returns the hash
func (obj *loopHeader) Hash() hash.Hash {
	return obj.hash
}

// IsCounter returns true if there is a counter, false otherwise
func (obj *loopHeader) IsCounter() bool {
	return obj.counter != nil
}

// Counter returns the counter
func (obj *loopHeader) Counter() LoopCounter {
	return obj.counter
}

// IsKeyValue returns true if there is a keyValue, false otherwise
func (obj *loopHeader) IsKeyValue() bool {
	return obj.keyValue != nil
}

// KeyValue returns the keyValue
func (obj *loopHeader) KeyValue() LoopKeyValue {
	return obj.keyValue
}

// IsInfinite returns true if infinite, false otherwise
func (obj *loopHeader) IsInfinite() bool {
	return obj.isInfinite
}
