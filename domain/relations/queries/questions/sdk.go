package questions

import "github.com/steve-care-software/steve/domain/relations/data/connections/links"

// Builder represents the question builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithLink(link links.Link) Builder
	Now() (Question, error)
}

// Question represents the question
type Question interface {
	Name() string
	Link() links.Link
}
