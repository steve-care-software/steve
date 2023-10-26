package inputs

import "github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links"

// Builder represents an input builder
type Builder interface {
	Create() Builder
	WithBytes(bytes []byte) Builder
	WithLink(link links.Link) Builder
	Now() (Input, error)
}

// Input represents an input
type Input interface {
	Bytes() []byte
	Link() links.Link
}
