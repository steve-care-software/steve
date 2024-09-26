package numerics

import (
	"github.com/steve-care-software/steve/domain/hash"
)

const (
	// FlagUint represents the uint flag
	FlagUint (uint8) = iota

	// FlagInt represents the int flag
	FlagInt

	// FlagFloat represents the float flag
	FlagFloat
)

const (
	// Size8 represents the size 8
	Size8 (uint8) = iota

	// Size16 represents the size 16
	Size16

	// Size32 represents the size 32
	Size32

	// Size64 represents the size 64
	Size64
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the numeric builder
type Builder interface {
	Create() Builder
	WithFlag(flag uint8) Builder
	WithSize(size uint8) Builder
	Now() (Numeric, error)
}

// Numeric represents a numeric kind
type Numeric interface {
	Hash() hash.Hash
	Flag() uint8 // uint, int, float
	Size() uint8 // 8, 16, 32, 64
}
