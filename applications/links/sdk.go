package links

import (
	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/connections/links"
)

// Application repreents a link application
type Application interface {
	Insert(link links.Link) error
	Delete(identifier uuid.UUID) error
}
