package kinds

import (
	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/containers/kinds/numerics"
)

type kind struct {
	hash       hash.Hash
	numeric    numerics.Numeric
	pEngine    *uint8
	pRemaining *uint8
}

func createKindWithNumeric(
	hash hash.Hash,
	numeric numerics.Numeric,
) Kind {
	return createKindInternally(
		hash,
		numeric,
		nil,
		nil,
	)
}

func createKindWithEngine(
	hash hash.Hash,
	pEngine *uint8,
) Kind {
	return createKindInternally(
		hash,
		nil,
		pEngine,
		nil,
	)
}

func createKindWithRemaining(
	hash hash.Hash,
	pRemaining *uint8,
) Kind {
	return createKindInternally(
		hash,
		nil,
		nil,
		pRemaining,
	)
}

func createKindInternally(
	hash hash.Hash,
	numeric numerics.Numeric,
	pEngine *uint8,
	pRemaining *uint8,
) Kind {
	out := kind{
		hash:       hash,
		numeric:    numeric,
		pEngine:    pEngine,
		pRemaining: pRemaining,
	}

	return &out
}

// Hash returns the hash
func (obj *kind) Hash() hash.Hash {
	return obj.hash
}

// IsNumeric returns true if there is numeric, false otherwise
func (obj *kind) IsNumeric() bool {
	return obj.numeric != nil
}

// IsNumeric returns true if there is numeric, false otherwise
func (obj *kind) Numeric() numerics.Numeric {
	return obj.numeric
}

// IsEngine returns true if there is an engine, false otherwise
func (obj *kind) IsEngine() bool {
	return obj.pEngine != nil
}

// Engine returns the engine, if any
func (obj *kind) Engine() *uint8 {
	return obj.pEngine
}

// IsRemaining returns true if there is a remaining, false otherwise
func (obj *kind) IsRemaining() bool {
	return obj.pRemaining != nil
}

// Remaining returns the remaining, if any
func (obj *kind) Remaining() *uint8 {
	return obj.pRemaining
}
