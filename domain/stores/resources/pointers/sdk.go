package pointers

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// Uint64Size represents the uint64 byte size
const Uint64Size = 8

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	builder := NewBuilder()
	pointerBuilder := NewPointerBuilder()
	return createAdapter(
		builder,
		pointerBuilder,
	)
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

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewPointerBuilder creates a new pointer builder
func NewPointerBuilder() PointerBuilder {
	return createPointerBuilder()
}

// Adapter represents a pointer adapter
type Adapter interface {
	InstancesToBytes(ins Pointers) ([]byte, error)
	BytesToInstances(data []byte) (Pointers, []byte, error)
	InstanceToBytes(ins Pointer) ([]byte, error)
	BytesToInstance(data []byte) (Pointer, []byte, error)
}

// Builder represents the pointers builder
type Builder interface {
	Create() Builder
	WithList(list []Pointer) Builder
	Now() (Pointers, error)
}

// Pointers represents pointers
type Pointers interface {
	List() []Pointer
}

// PointerBuilder represents the pointer builder
type PointerBuilder interface {
	Create() PointerBuilder
	WithIndex(index uint) PointerBuilder
	WithLength(length uint) PointerBuilder
	Now() (Pointer, error)
}

// Pointer represents a pointer
type Pointer interface {
	Index() uint
	Length() uint
}
