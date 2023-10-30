package headers

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources/headers/identifiers"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources/headers/links"
)

// Builder represents an header builder
type Builder interface {
	Create() Builder
	WithIdentifier(identifier identifiers.Identifier) Builder
	WithLink(link links.Link) Builder
	Now() (Header, error)
}

// Header represents a resource header
type Header interface {
	Identifier() identifiers.Identifier
	Link() links.Link
}
