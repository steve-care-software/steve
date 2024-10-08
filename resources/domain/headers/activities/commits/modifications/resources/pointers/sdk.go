package pointers

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/hash"
)

// Uint64Size represents the uint64 byte size
const Uint64Size = 8

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	builder := NewBuilder()
	return createAdapter(
		builder,
	)
}

// NewBuilder creates a new pointer builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents a pointer adapter
type Adapter interface {
	ToBytes(ins Pointer) ([]byte, error)
	ToInstance(data []byte) (Pointer, []byte, error)
}

// Builder represents the pointer builder
type Builder interface {
	Create() Builder
	WithIndex(index uint) Builder
	WithLength(length uint) Builder
	Now() (Pointer, error)
}

// Pointer represents a pointer
type Pointer interface {
	Hash() hash.Hash
	Index() uint
	Length() uint
}

// Uint64ToBytes converts uint64 to bytes
func Uint64ToBytes(value uint64) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, value)
	return bytes
}

// BytesToUint64 converts bytes to uint64
func BytesToUint64(data []byte) (*uint64, error) {
	if len(data) != Uint64Size {
		str := fmt.Sprintf("byte slice must be exactly %d bytes long to convert to uint64", Uint64Size)
		return nil, errors.New(str)
	}

	value := binary.BigEndian.Uint64(data)
	return &value, nil
}
