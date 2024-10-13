package updates

import (
	"github.com/steve-care-software/steve/hash"
)

const (
	// KindReference represents the reference kind
	KindReference (uint8) = iota
)

// UpdateBuilder represents an update builder
type UpdateBuilder interface {
	Create() UpdateBuilder
	WithKind(kind uint8) UpdateBuilder
	WithHashes(hashes []hash.Hash) UpdateBuilder
	Now() (Update, error)
}

// Update represents a database update
type Update interface {
	Kind() uint8
	Hashes() []hash.Hash
}
