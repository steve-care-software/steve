package uuids

import "github.com/google/uuid"

// UUIDSize represents the uuid size
const UUIDSize = 16

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	return createAdapter()
}

// Adapter represents the uuid adapyter
type Adapter interface {
	FromBytes(data []byte) ([]uuid.UUID, error)
	FromInstances(list []uuid.UUID) ([]byte, error)
}
