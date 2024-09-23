package references

import "github.com/steve-care-software/steve/domain/hash"

// Reference represents a reference
type Reference interface {
	Hash() hash.Hash
	Grammar() string
	Block() string
}
