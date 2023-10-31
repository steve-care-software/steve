package deletes

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources/headers/identifiers"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources/headers/links"
)

// Builder represents a delete builder
type Builder interface {
	Create() Builder
	WithLink(link links.Link) Builder
	WithIdentifier(identifier identifiers.Identifier) Builder
	Now() (Delete, error)
}

// Delete represents a delete
type Delete interface {
	IsLink() bool
	Link() links.Link
	IsIdentifier() bool
	Identifier() identifiers.Identifier
}
