package uuids

import "github.com/google/uuid"

// Adapter represents the uuid adapyter
type Adapter interface {
	FromBytes(data []byte) ([]uuid.UUID, error)
	FromInstances(list []uuid.UUID) ([]byte, error)
}
