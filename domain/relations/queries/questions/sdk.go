package questions

import "github.com/steve-care-software/steve/domain/relations/data/connections/links"

// Question represents the question
type Question interface {
	Name() string
	Link() links.Link
}
