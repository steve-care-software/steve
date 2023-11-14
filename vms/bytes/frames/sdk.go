package frames

import (
	"github.com/steve-care-software/steve/vms/bytes/results/hash"
)

// Frames represents frames
type Frames interface {
	Hash() hash.Hash
	List() []Frame
}

// Frame represents a frame
type Frame interface {
	Hash() hash.Hash
	Fetch(name string) ([]byte, error)
}
