package inserts

import "github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources"

// Builder represents an inserts builder
type Builder interface {
	Create() Builder
	WithList(list []Insert) Builder
	Now() (Inserts, error)
}

// Inserts represents inserts
type Inserts interface {
	List() []Insert
}

// InsertBuilder represents an insert builder
type InsertBuilder interface {
	Create() InsertBuilder
	WithResource(resource resources.Resource) InsertBuilder
	WithBytes(bytes []byte) InsertBuilder
	Now() (Insert, error)
}

// Insert represents an insert
type Insert interface {
	Resource() resources.Resource
	Bytes() []byte
}
